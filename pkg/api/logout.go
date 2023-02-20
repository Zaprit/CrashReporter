package api

import (
	"net/http"

	"github.com/Zaprit/CrashReporter/pkg/db"
	"github.com/gin-gonic/gin"
)

func LogoutHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		sessionId, err := context.Cookie("session_id")
		if err != nil {
			context.String(http.StatusBadRequest, "User not logged in.")
			return
		}
		err = db.EndSession(sessionId)
		if err != nil {
			context.String(http.StatusBadRequest, "Session not found.")
			return
		}
		context.Redirect(http.StatusTemporaryRedirect, "/")
	}
}
