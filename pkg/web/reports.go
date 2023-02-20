package web

import (
	"github.com/Zaprit/CrashReporter/pkg/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReportsHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		sessionID, err := context.Cookie("session_id")
		if err != nil {
			context.HTML(http.StatusUnauthorized, "unauthorized.gohtml", nil)
			return
		}
		session, err := db.GetSession(sessionID)
		if err != nil {
			context.HTML(http.StatusUnauthorized, "unauthorized.gohtml", nil)
			return
		}

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
			"Notices":  db.GetNotifications(),
			"Username": session.Username,
			"Avatar":   session.AvatarURL,
			"Reports":  reports,
		})
	}
}
