package web

import (
	"github.com/Zaprit/CrashReporter/pkg/db"
	"github.com/Zaprit/CrashReporter/pkg/lighthouse_client"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminReportHandler() gin.HandlerFunc {
	return func(context *gin.Context) {

		userName := context.GetString("session_user")
		avatarURL := context.GetString("session_avatar")

		report := db.GetReport(context.Query("id"))

		if report.UUID == "" {
			context.HTML(404, "not_found.gohtml", nil)
			return
		}

		reportAvatar := "/static/missing.png"

		user, _, err := lighthouse_client.GetUser(report.Username)
		if err == nil {
			reportAvatar = lighthouse_client.UserAvatar(user)
		}

		context.HTML(http.StatusOK, "admin_report.gohtml", gin.H{
			"LoggedIn":          userName != "",
			"Username":          userName,
			"Avatar":            avatarURL,
			"AdminArea":         true,
			"Notices":           db.GetNotifications(),
			"ReportTitle":       report.Title,
			"ReportUUID":        report.UUID,
			"ReportUser":        report.Username,
			"ReportAvatar":      reportAvatar,
			"ReportPlatform":    report.Platform,
			"ReportTime":        report.SubmitTime,
			"ReportIP":          report.SubmitterIP,
			"ReportType":        report.Type,
			"ReportDescription": report.Description,
			"ReportEvidence":    report.Evidence,
			"ReportResolved":    report.Resolved,
		})
	}
}
