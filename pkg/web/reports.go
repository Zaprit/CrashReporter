package web

import (
	"github.com/Zaprit/CrashReporter/pkg/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReportsHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		userName := context.GetString("session_user")
		avatarURL := context.GetString("session_avatar")
		// TODO: actually write this later
		//var skip int
		//
		//if context.Query("p") != "" {
		//    page, err := strconv.ParseInt(context.Query("p"), 10, 16)
		//    if err != nil {
		//        skip = 0
		//    } else {
		//        skip = int(page)*10
		//    }
		//}

		reports := db.GetReports()

		context.HTML(http.StatusOK, "reports.gohtml", gin.H{
			"LoggedIn":  userName != "",
			"Username":  userName,
			"Avatar":    avatarURL,
			"AdminArea": true,
			"Notices":   db.GetNotifications(),
			"Reports":   reports,
		})
	}
}
