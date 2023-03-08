package main

import (
	"github.com/Zaprit/CrashReporter/pkg/api"
	"github.com/Zaprit/CrashReporter/pkg/config"
	"github.com/Zaprit/CrashReporter/pkg/db"
	"github.com/Zaprit/CrashReporter/pkg/importer"
	"github.com/Zaprit/CrashReporter/pkg/middleware"
	"github.com/Zaprit/CrashReporter/pkg/web"
	"github.com/gin-gonic/gin"
	"io/fs"
	"log"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		if pathErr, ok := err.(*fs.PathError); ok {
			config.LoadedConfig = config.DefaultConfig
			err := config.SaveConfig()
			if err != nil {
				log.Fatalln(err.Error())
			}
		} else {
			log.Fatalln(pathErr.Error())
		}
	}

	db.OpenDB(config.LoadedConfig.DBFile)
	db.MigrateDB()

	if config.LoadedConfig.MigrateOld {
		importer.ImportOldReports()
		config.LoadedConfig.MigrateOld = false
		err := config.SaveConfig()
		if err != nil {
			log.Println("Error In Migrating Reports ", err.Error())
		}
	}

	router := gin.Default()
	err = router.SetTrustedProxies(nil)
	if err != nil {
		panic(err.Error())
	}

	router.LoadHTMLGlob("static/partials/*")

	router.StaticFile("/favicon.png", "static/img/CrashHelper.png")

	router.Static("/static/styles", "static/styles")
	router.Static("/static/img", "static/img")
	router.Static("/static/js", "static/js")

	router.Use(middleware.SessionMiddleware())

	publicPages := router.Group("/", middleware.SessionMiddleware())
	publicPages.GET("/", web.IndexHandler())
	publicPages.GET("/login", web.LoginHandler())
	publicPages.GET("/report", web.ReportHandler())

	adminPages := router.Group("/admin", middleware.SessionMiddleware(), middleware.AuthorizationMiddleware())
	adminPages.GET("/report", web.ReportHandler())
	adminPages.GET("/reports", web.ReportsHandler())
	adminPages.GET("/", web.AdminDashboardHandler())

	apiV1 := router.Group("/api/v1/")

	apiV1.POST("report", api.SubmitReportHandler())
	apiV1.GET("oauth/callback", api.OAuthCallbackHandler())
	apiV1.GET("user/:username", api.LighthouseUsersApiHandler())
	apiV1.GET("search", api.LighthouseUserSearchApiHandler())
	apiV1.GET("report/:uuid/comments", api.CommentsHandler())
	apiV1.POST("report/:uuid/post_comment", api.PostCommentHandler())
	apiV1.GET("logout", api.LogoutHandler())

	router.NoRoute(func(context *gin.Context) {
		context.HTML(404, "not_found.gohtml", nil)
	})

	err = router.Run(config.LoadedConfig.ListenAddress)
	if err != nil {
		panic(err)
	}
}
