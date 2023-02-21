package web

import (
	"github.com/Zaprit/CrashReporter/pkg/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		userName := context.GetString("session_user")
		avatarURL := context.GetString("session_avatar")

		context.HTML(http.StatusOK, "index.gohtml", gin.H{
			"Notices":          db.GetNotifications(),
			"ReportCategories": db.GetReportCategories(),
			"LoggedIn":         userName != "",
			"Username":         userName,
			"Avatar":           avatarURL,
		})
	}
}
