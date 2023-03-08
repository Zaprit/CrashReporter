package github_oauth

import (
	"encoding/json"
	"github.com/Zaprit/CrashReporter/pkg/db"
	"github.com/Zaprit/CrashReporter/pkg/model"
	"io"
	"log"
	"net/http"
)

// UserFromOAuth2Token gets a session from the auth token we receive from GitHub
func UserFromOAuth2Token(token string) (model.Session, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return model.Session{}, err
	}

	req.Header.Add("Authorization", "Bearer "+token)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in response.\n[ERROR] -", err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	var session model.Session
	err = json.Unmarshal(body, &session)
	if err != nil {
		return model.Session{}, err
	}

	session.OAuthToken = token

	err = db.SaveSession(&session)

	return session, err
}
