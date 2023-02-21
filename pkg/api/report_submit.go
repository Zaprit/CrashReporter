package api

import (
	"net/http"
	"time"

	"github.com/Zaprit/CrashReporter/pkg/db"
	"github.com/Zaprit/CrashReporter/pkg/lighthouse_client"
	"github.com/Zaprit/CrashReporter/pkg/model"
	"github.com/Zaprit/CrashReporter/pkg/webhook"
	"github.com/gin-gonic/gin"
)

func SubmitReportHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		report := model.Report{
			Title:       context.PostForm("title"),
			Username:    context.PostForm("username"),
			Type:        context.PostForm("issuetype"),
			Platform:    context.PostForm("platform"),
			Description: context.PostForm("details"),
			SubmitterIP: context.RemoteIP(),
			SubmitTime:  time.Now().UTC(),
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

		if report.Type == "LBP issue, i.e. game crash, graphics bugs, etc." {
			context.String(http.StatusUnprocessableEntity, "This service is not for submitting reports regarding LittleBigPlanet. Please only submit reports pertaining to Beacon.<br />If you require immediate assistance, we recommend asking for help in the <span class='branch-name'>🚀union-space-corps</span> channel of the <a href='https://discord.gg/lbpunion'>LBP Union Discord</a>.")
		}

		if !db.ReportTypeExists(report.Type) {
			context.String(http.StatusBadRequest, "Invalid Type, this is most likely an error, please contact an administrator")
			return
		}

		if context.PostForm("hasevidence") == "on" {
			report.Evidence = true
		}

		err = db.SubmitReport(&report)
		if err != nil {
			context.String(http.StatusInternalServerError, "Failed to submit report, please contact an administrator")
			return
		}

		go webhook.Sendreport(report)

		context.String(http.StatusCreated, "Your report has been submitted, your report ID is <code>%s</code>", report.UUID)

	}
}
