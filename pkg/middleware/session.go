package middleware

import (
	"github.com/Zaprit/CrashReporter/pkg/db"
	"github.com/gin-gonic/gin"
)

func SessionMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		sessionId, err := context.Cookie("session_id")
		if err != nil {
			return
		}
		context.Set("session_id", sessionId)

		session, _ := db.GetSession(sessionId)

		context.Set("session_user", session.Username)
		context.Set("session_avatar", session.AvatarURL)

		context.Next()
	}
}
