package web

import (
    "github.com/Zaprit/CrashReporter/pkg/db"
    "github.com/gin-gonic/gin"
    "net/http"
)

func IndexHandler() gin.HandlerFunc {
    return func(context *gin.Context) {

        context.HTML(http.StatusOK, "index.gohtml", gin.H{
            "Notices": db.GetNotifications(),
            "ReportCategories": db.GetReportCategories(),
        })
    }
}