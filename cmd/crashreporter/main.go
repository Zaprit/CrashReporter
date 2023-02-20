package main

import (
	"log"

	"github.com/Zaprit/CrashReporter/pkg/api"
	"github.com/Zaprit/CrashReporter/pkg/config"
	"github.com/Zaprit/CrashReporter/pkg/db"
	"github.com/Zaprit/CrashReporter/pkg/web"
	"github.com/gin-gonic/gin"
)

func main() {
	err := config.LoadConfig("config.toml")
	if err != nil {
		log.Fatalln(err.Error())
	}

	db.OpenDB(config.LoadedConfig.DBFile)
	db.MigrateDB()

	router := gin.Default()
	err = router.SetTrustedProxies(nil)
	if err != nil {
		panic(err.Error())
	}

	router.LoadHTMLGlob("static/partials/*")

	router.Static("/static/styles", "static/styles")
	router.Static("/static/js", "static/js")

	router.GET("/", web.IndexHandler())
	router.GET("/admin/report", web.AdminReportHandler())

	router.GET("/admin/reports", web.ReportsHandler())

	router.POST("/api/v1/report", api.SubmitReportHandler())
	router.GET("/api/v1/oauth/callback", api.OAuthCallbackHandler())
	router.GET("/api/v1/user/:username", api.LighthouseUsersApiHandler())
	router.GET("/api/v1/report/:uuid/comments", api.CommentsHandler())
	router.GET("/api/v1/logout", api.LogoutHandler())

	err = router.Run(config.LoadedConfig.ListenAddress)
	if err != nil {
		panic(err)
	}
}
