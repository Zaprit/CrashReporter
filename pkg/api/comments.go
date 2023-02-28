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
		uuid := context.Param("uuid")

		report, err := db.GetReportID(uuid)
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
		uuid := context.Param("uuid")

		user := context.GetString("session_user")
		avatar := context.GetString("session_avatar")

		id, err := db.GetReportID(uuid)
		if err != nil {
			context.String(http.StatusNotFound, "Report does not exist")
			return
		}

		comment := model.Comment{
			Poster:       user,
			PosterAvatar: avatar,
			ReportID:     id,
			CreateTime:   time.Now(),
			Content:      context.PostForm("content"),
		}

		db.PostComment(comment)

		context.String(http.StatusOK, "Comment Posted")
	}
}
