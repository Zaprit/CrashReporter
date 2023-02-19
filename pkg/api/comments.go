package api

import (
    "github.com/Zaprit/CrashReporter/pkg/db"
    "github.com/gin-gonic/gin"
    "net/http"
)

type ReportURI struct {
    ReportUUID string `uri:"uuid" binding:"required"`
}

func CommentsHandler() gin.HandlerFunc {
    return func(context *gin.Context) {
        var uridata ReportURI

        _ = context.BindUri(&uridata)

        report := db.GetReport(uridata.ReportUUID, true)

        if report.ID == 0 {
            context.String(http.StatusNotFound, "Report Not Found")
            return
        }

        context.JSON(http.StatusOK, report.Comments)

    }
}