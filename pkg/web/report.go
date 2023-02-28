package web

import (
	"github.com/Zaprit/CrashReporter/pkg/db"
	"github.com/Zaprit/CrashReporter/pkg/lighthouse_client"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReportHandler() gin.HandlerFunc {
	return func(context *gin.Context) {

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

		context.HTML(http.StatusOK, "report.gohtml", gin.H{
			"Notices":           db.GetNotifications(),
			"ReportTitle":       report.Title,
			"ReportUUID":        report.UUID,
			"ReportUser":        report.Username,
			"ReportPlatform":    report.Platform,
			"ReportAvatar":      reportAvatar,
			"ReportTime":        report.SubmitTime,
			"ReportType":        report.Type,
			"ReportDescription": report.Description,
			"ReportEvidence":    report.Evidence,
			"ReportPriority":    report.Priority,
			"ReportResolved":    report.Resolved,
		})
	}
}
