package middleware

import (
	"github.com/Zaprit/CrashReporter/pkg/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func APIAuthorizationMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		sessionId, _ := context.Cookie("session_id")
		context.Set("session_id", sessionId)

		session, _ := db.GetSession(sessionId)

		if session.ID == "" {
			context.String(http.StatusForbidden, "Unauthorized")
			context.Abort()
			return
		}
		context.Next()
	}
}
