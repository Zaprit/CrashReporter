package web

import (
	"fmt"
	"github.com/Zaprit/CrashReporter/pkg/config"
	"github.com/Zaprit/CrashReporter/pkg/db"
	"github.com/Zaprit/CrashReporter/pkg/lighthouse_client"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Truncate(text string, width int) string {

	r := []rune(text)
	trunc := r[:width]
	return string(trunc) + "..."
}
func ReportHandler() gin.HandlerFunc {
	return func(context *gin.Context) {

		report := db.GetReport(context.Query("id"))

		if report.UUID == "" {
			fmt.Println("Report Not Found")
			context.HTML(404, "not_found.gohtml", nil)
			return
		}

		reportAvatar := "/static/missing.png"

		user, _, err := lighthouse_client.GetUser(report.Username)
		if err == nil {
			reportAvatar = lighthouse_client.UserAvatar(user)
		}

		summary := Truncate(report.Description, 50)
		fmt.Println(report.Description)
		context.HTML(http.StatusOK, "report.gohtml", gin.H{
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
		})
	}
}
