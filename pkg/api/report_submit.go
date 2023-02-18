package api

import (
    "github.com/Zaprit/CrashReporter/pkg/model"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "net/http"
)

func SubmitReportHandler() gin.HandlerFunc {
    return func(context *gin.Context) {
        report := model.Report{
            Title:       context.PostForm("title"),
            Username:    context.PostForm("username"),
            Type:        context.PostForm("issuetype"),
            Platform:    context.PostForm("platform"),
            Description: context.PostForm("details"),
            Evidence:    false,
        }

        if len(report.Title) > 20 {
            context.String(http.StatusBadRequest, "The title is too long, it must be less than 20 characters.")
        }

        if len(report.Username) > 16 {
            context.String(http.StatusBadRequest, "The username is too long, it must be less than 16 characters.")
        }

        if len(report.Description) > 500 {
            context.String(http.StatusBadRequest, "The description is too long, it must be less than or equal to 500 characters in length")
        }

        _, platformExists := model.ValidPlatforms[report.Platform]
        if !platformExists {
            context.String(http.StatusBadRequest, "Invalid platform, this is most likely an error, please report in #union-space-corps")
        }

        if


        if context.PostForm("hasevidence") == "on" {
            report.Evidence = true
        }

    }
}