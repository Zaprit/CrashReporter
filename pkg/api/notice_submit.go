package api

import (
	"github.com/Zaprit/CrashReporter/pkg/db"
	"github.com/Zaprit/CrashReporter/pkg/model"
	"github.com/gin-gonic/gin"
	"time"
)

func NoticeSubmitHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		notice := model.Notice{
			Status:     context.PostForm("status"),
			NoticeText: context.PostForm("content"),
			Date:       time.Now(),
			Ended:      false,
		}

		db.PostNotice(notice)
	}
}
