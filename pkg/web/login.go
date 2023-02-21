package web

import (
	"github.com/Zaprit/CrashReporter/pkg/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		userName := context.GetString("session_user")
		if userName != "" {
			context.Redirect(http.StatusFound, "/admin/reports")
			return
		}

		context.HTML(http.StatusOK, "login.gohtml", gin.H{
			"ClientID": config.LoadedConfig.OAuth2ClientID,
			"State":    "e73fb15232e3b4bbee2517451ac5e1dd",
		})
	}
}
