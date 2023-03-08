package api

import (
	"encoding/json"
	"github.com/Zaprit/CrashReporter/pkg/config"
	"github.com/Zaprit/CrashReporter/pkg/github_oauth"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func OAuthCallbackHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		code := context.Query("code")
		if code == "" {
			context.String(http.StatusBadRequest, "Invalid code")
			return
		}

		values := url.Values{
			"code":          {code},
			"client_id":     {config.LoadedConfig.OAuth2ClientID},
			"client_secret": {config.LoadedConfig.OAuth2ClientSecret},
		}
		req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", strings.NewReader(values.Encode()))
		if err != nil {
			context.String(http.StatusInternalServerError, "Failed to contact the oauth provider")
			return
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			context.String(http.StatusInternalServerError, "Failed to contact the oauth provider.")
			return
		}

		if resp.StatusCode != http.StatusOK {
			context.HTML(http.StatusInternalServerError, "login_error.html", gin.H{
				"ErrorDescription": "Failed to contact the OAuth2 Provider.",
			})
			return
		}

		var Token struct {
			AccessToken string `json:"access_token"`
			Scope       string `json:"scope"`
			TokenType   string `json:"token_type"`
		}
		jsonBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			context.String(http.StatusInternalServerError, "Failed to parse response from github")
			return
		}

		err = json.Unmarshal(jsonBytes, &Token)
		if err != nil {
			context.String(http.StatusInternalServerError, "Failed to parse response from github.")
			return
		}

		session, err := github_oauth.UserFromOAuth2Token(Token.AccessToken)

		var isAdmin bool
		for _, admin := range config.LoadedConfig.AdminUsers {
			if session.Username == admin {

				isAdmin = true
			}
		}

		if !isAdmin {
			context.HTML(http.StatusUnauthorized, "unauthorized.gohtml", nil)
			return
		}

		context.SetCookie("session_id", session.ID, 3600, "/", config.LoadedConfig.PublicURL, true, true)

		context.Redirect(http.StatusFound, "/admin/reports")
	}
}
