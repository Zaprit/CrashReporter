package web

import (
	"github.com/Zaprit/CrashReporter/pkg/db"
	"github.com/Zaprit/CrashReporter/pkg/lighthouse_client"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminDashboardHandler() gin.HandlerFunc {
	return func(context *gin.Context) {

		userName := context.GetString("session_user")
		avatarURL := context.GetString("session_avatar")

		reports := db.GetReports()

		for i, r := range reports {
			reportAvatar := "/static/missing.png"

			user, _, err := lighthouse_client.GetUser(r.Username)
			if err == nil {
				reportAvatar = lighthouse_client.UserAvatar(user)
			}
			reports[i].Avatar = reportAvatar
		}

		context.HTML(http.StatusOK, "admin_dashboard.gohtml", gin.H{
			"LoggedIn":  userName != "",
			"Username":  userName,
			"Avatar":    avatarURL,
			"AdminArea": true,
			"Notices":   db.GetNotifications(),
			"Reports":   reports,
		})
	}
}
