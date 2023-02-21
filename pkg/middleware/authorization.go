package middleware

import (
	"github.com/Zaprit/CrashReporter/pkg/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		sessionId, _ := context.Cookie("session_id")
		context.Set("session_id", sessionId)

		session, _ := db.GetSession(sessionId)

		if session.ID == "" {
			context.HTML(http.StatusUnauthorized, "unauthorized.gohtml", nil)
			context.Abort()
			return
		}
		context.Next()
	}
}
