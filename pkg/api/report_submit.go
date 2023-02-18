package api

import (
    "fmt"
    "github.com/Zaprit/CrashReporter/pkg/db"
    "github.com/Zaprit/CrashReporter/pkg/model"
    "github.com/gin-gonic/gin"
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

        crap, _ := context.GetRawData()

        fmt.Println(string(crap))

        if len(report.Title) > 20 {
            context.String(http.StatusBadRequest, "The title is too long, it must be less than 20 characters.")
            return
        }

        if len(report.Username) > 16 {
            context.String(http.StatusBadRequest, "The username is too long, it must be less than 16 characters.")
            return
        }

        if len(report.Description) > 500 {
            context.String(http.StatusBadRequest, "The description is too long, it must be less than or equal to 500 characters in length")
            return
        }

        _, platformExists := model.ValidPlatforms[report.Platform]
        if !platformExists {
            context.String(http.StatusBadRequest, "Invalid platform, this is most likely an error, please report in #union-space-corps")
            return
        }

        fmt.Printf("'%s'",report.Type)
        if report.Type == "Choose an option" {
            context.String(http.StatusBadRequest, "Please choose an issue type")
            return
        }

        if !db.ReportTypeExists(report.Type) {
            context.String(http.StatusBadRequest, "Invalid Type, this is most likely an error, please contact an administrator")
            return
        }


        if context.PostForm("hasevidence") == "on" {
            report.Evidence = true
        }

        err := db.SubmitReport(&report)
        if err != nil {
            context.String(http.StatusInternalServerError, "Failed to submit report, please contact an administrator")
            return
        }
        context.String(http.StatusCreated, "Your report has been submitted, your report ID is <code>%s</code>", report.ID)

    }
}