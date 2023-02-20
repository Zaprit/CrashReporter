package web

import (
	"github.com/Zaprit/CrashReporter/pkg/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReportHandler() gin.HandlerFunc {
	return func(context *gin.Context) {

		context.HTML(http.StatusOK, "report.gohtml", gin.H{
			"Notices": db.GetNotifications(),
			"Report":  db.GetReport(context.Query("id"), true),
		})
	}
}
