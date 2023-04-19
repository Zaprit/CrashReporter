package web

import (
	"github.com/Zaprit/CrashReporter/pkg/config"
	"github.com/Zaprit/CrashReporter/pkg/db"
	"github.com/Zaprit/CrashReporter/pkg/lighthouse_client"
	"github.com/gin-gonic/gin"
	"net/http"
)

func truncate(text string, width int) string {
	if len(text) <= width {
		return text
	}

	r := []rune(text)
	trunc := r[:width]
	return string(trunc) + "..."
}

func ReportHandler() gin.HandlerFunc {
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

		if userName != "" {
			db.ReadReport(report.UUID)
		}

		summary := truncate(report.Description, 50)

		context.HTML(http.StatusOK, "report.gohtml", gin.H{
			"AdminArea":         context.FullPath() == "/admin/report",
			"LoggedIn":          userName != "",
			"Username":          userName,
			"Avatar":            avatarURL,
			"Notices":           db.GetNotifications(),
			"LighthouseURL":     config.LoadedConfig.LighthouseURL,
			"ReportTitle":       report.Title,
			"ReportUUID":        report.UUID,
			"ReportUser":        report.Username,
			"ReportPlatform":    report.Platform,
			"ReportAvatar":      reportAvatar,
			"ReportTime":        report.SubmitTime,
			"ReportType":        report.Type,
			"ReportDescription": report.Description,
			"ReportSummary":     summary,
			"ReportEvidence":    report.Evidence,
			"ReportPriority":    report.Priority,
			"ReportResolved":    report.Resolved,
			"ReportIP":          report.SubmitterIP,
		})
	}
}
