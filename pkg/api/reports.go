package api

import (
	"github.com/Zaprit/CrashReporter/pkg/db"
	"github.com/Zaprit/CrashReporter/pkg/lighthouse_client"
	"github.com/Zaprit/CrashReporter/pkg/model"
	"github.com/Zaprit/CrashReporter/pkg/webhook"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func SubmitReportHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		report := model.Report{
			Title:       context.PostForm("title"),
			Username:    context.PostForm("username"),
			Type:        context.PostForm("issue_type"),
			Platform:    context.PostForm("platform"),
			Description: context.PostForm("details"),
			SubmitterIP: context.RemoteIP(),
			Evidence:    false,
		}

		if len(report.Title) > 20 {
			context.String(http.StatusBadRequest, "The title is too long, it must be less than 20 characters.")
			return
		}

		if len(report.Username) < 3 {
			context.String(http.StatusBadRequest, "The username is too short, it must be at least 3 characters.")
			return
		}

		if len(report.Username) > 16 {
			context.String(http.StatusBadRequest, "The username is too long, it must be less than 16 characters.")
			return
		}

		user, _, err := lighthouse_client.GetUser(report.Username)
		if err != nil || user.UserId == 0 {
			context.String(http.StatusBadRequest, "Lighthouse user doesn't exist, please enter a valid username")
			return
		}

		report.UserID = user.UserId

		if len(report.Description) > 500 {
			context.String(http.StatusBadRequest, "The description is too long, it must be less than or equal to 500 characters in length")
			return
		}

		_, platformExists := model.ValidPlatforms[report.Platform]
		if !platformExists {
			context.String(http.StatusBadRequest, "Invalid platform, this is most likely an error, please report in #union-space-corps")
			return
		}

		if report.Type == "Choose an option" {
			context.String(http.StatusBadRequest, "Please choose an issue type")
			return
		}

		if report.Type == "LBP issue, i.e. game crash, graphics bugs, etc." { //TODO: Move all the names from long UI strings to more generic values
			context.String(http.StatusUnprocessableEntity, "This service is not for submitting reports regarding LittleBigPlanet. Please only submit reports pertaining to Beacon.<br />If you require immediate assistance, we recommend asking for help in the <span class='branch-name'>🚀union-space-corps</span> channel of the <a href='https://discord.gg/lbpunion'>LBP Union Discord</a>.")
		}

		if context.PostForm("priority") == "" {
			context.String(http.StatusBadRequest, "Please choose a priority")
		}

		if priority, exists := model.Priorities[context.PostForm("priority")]; exists {
			report.Priority = priority
		} else {
			context.String(http.StatusBadRequest, "Invalid Priority")
		}

		if !db.ReportTypeExists(report.Type) {
			context.String(http.StatusBadRequest, "Invalid Type, this is most likely an error, please contact an administrator")
			return
		}

		if context.PostForm("has_evidence") == "on" {
			report.Evidence = true
		}

		err = db.SaveReport(&report)
		if err != nil {
			context.String(http.StatusInternalServerError, "Failed to submit report, please contact an administrator")
			return
		}

		go func() {
			id, err := webhook.SendReport(report)
			if err != nil {
				log.Println("failed to send discord embed: ", err.Error())
			}
			report.DiscordMessageID = id
			err = db.SaveReport(&report)
			if err != nil {
				log.Println("failed to save discord message id: ", err.Error())
			}
		}()

		context.String(http.StatusCreated, "Your report has been submitted, your report ID is <code>%s</code>", report.UUID)

	}
}

func ReportDismissHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		reportID := context.Param("uuid")
		if reportID == "" {
			context.String(http.StatusBadRequest, "Invalid ID")
		}
		err := db.DismissReport(reportID)
		if err != nil {
			context.String(http.StatusInternalServerError, "Failed to dismiss report")
		}
	}
}

func ReportOpenHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		reportID := context.Param("uuid")
		if reportID == "" {
			context.String(http.StatusBadRequest, "Invalid ID")
		}
		err := db.ReopenReport(reportID)
		if err != nil {
			context.String(http.StatusInternalServerError, "Failed to re-open report")
		}
	}
}
