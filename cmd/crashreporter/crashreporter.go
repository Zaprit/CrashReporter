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
	// Load Config
	err := config.LoadConfig()
	if err != nil {
		if _, ok := err.(*fs.PathError); ok {
			config.LoadedConfig = config.DefaultConfig
			err := config.SaveConfig()
			if err != nil {
				log.Fatalln(err.Error())
			}
		} else {
			log.Fatalln(err.Error())
		}
	}

	// Load DB
	switch config.LoadedConfig.Database.Type {
	case db.TypeSqlite:
		db.OpenSQLiteDB(config.LoadedConfig.Database.Database)
	case db.TypeMysql:
		db.OpenMySQLDB(
			config.LoadedConfig.Database.Hostname,
			config.LoadedConfig.Database.Port,
			config.LoadedConfig.Database.Database,
			config.LoadedConfig.Database.Username,
			config.LoadedConfig.Database.Password,
		)
	case db.TypePostgres:
		db.OpenPostgreSQLDB(
			config.LoadedConfig.Database.Hostname,
			config.LoadedConfig.Database.Port,
			config.LoadedConfig.Database.Database,
			config.LoadedConfig.Database.Username,
			config.LoadedConfig.Database.Password,
		)
	}

	db.MigrateDB()

	// Migrate old data
	if config.LoadedConfig.MigrateOld {
		importer.ImportOldReports()
		config.LoadedConfig.MigrateOld = false
		err := config.SaveConfig()
		if err != nil {
			log.Println("Error In Migrating Reports ", err.Error())
		}
	}

	if !config.LoadedConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	// HTTP Router
	router := gin.Default()
	router.TrustedPlatform = gin.PlatformCloudflare
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
	apiV1.GET("logout", api.LogoutHandler())
	authAPI := apiV1.Group("", middleware.APIAuthorizationMiddleware())
	authAPI.POST("notice", api.NoticeSubmitHandler())
	authAPI.DELETE("notice/:id", api.NoticeDismissHandler())

	authAPI.POST("report/:uuid/reopen", api.ReportOpenHandler())
	authAPI.DELETE("report/:uuid", api.ReportDismissHandler())
	authAPI.POST("report/:uuid/post_comment", api.PostCommentHandler())

	// 404 Page
	router.NoRoute(func(context *gin.Context) {
		context.HTML(404, "not_found.gohtml", nil)
	})

	// Run router
	err = router.Run(config.LoadedConfig.ListenAddress)
	if err != nil {
		panic(err)
	}
}
