package api

import (
	"github.com/Zaprit/CrashReporter/pkg/db"
	"github.com/Zaprit/CrashReporter/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ReportURI struct {
	ReportUUID string `uri:"uuid" binding:"required"`
}

func CommentsHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		var data ReportURI

		_ = context.BindUri(&data)

		report, err := db.GetReportID(data.ReportUUID)
		if err != nil {
			context.String(http.StatusNotFound, "Report Not Found")
			return
		}

		comments := db.GetComments(report)

		context.JSON(http.StatusOK, comments)
	}
}

func PostCommentHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		var data ReportURI
		_ = context.BindUri(&data)

		user := context.GetString("session_user")
		avatar := context.GetString("session_user")

		id, err := db.GetReportID(data.ReportUUID)
		if err != nil {
			context.Error(err)
			context.String(http.StatusNotFound, "Report does not exist")
			return
		}

		comment := model.Comment{
			Poster:       user,
			PosterAvatar: avatar,
			ReportID:     id,
			CreateTime:   time.Time{},
			Content:      context.PostForm("content"),
		}

		db.PostComment(comment)

		context.String(http.StatusOK, "Comment Posted")
	}
}
