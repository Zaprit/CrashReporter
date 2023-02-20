package web

import (
	"github.com/Zaprit/CrashReporter/pkg/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminReportHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		sessionID, err := context.Cookie("session_id")
		if err != nil {
			context.HTML(http.StatusUnauthorized, "unauthorized.gohtml", nil)
			return
		}
		session, err := db.GetSession(sessionID)
		if err != nil {
			context.HTML(http.StatusUnauthorized, "unauthorized.gohtml", nil)
			return
		}

		report := db.GetReport(context.Query("id"), true)

		context.HTML(http.StatusOK, "admin_report.gohtml", gin.H{
			"Notices":  db.GetNotifications(),
			"ReportID": report.UUID,
			"Username": session.Username,
			"Avatar":   session.AvatarURL,
		})
	}
}
