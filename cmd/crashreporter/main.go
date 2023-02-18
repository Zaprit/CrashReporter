package main

import (
    "github.com/Zaprit/CrashReporter/pkg/db"
    "github.com/Zaprit/CrashReporter/pkg/web"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

func main() {
    db.OpenDB("test.db")
    db.MigrateDB()

    uuid.EnableRandPool()

    router := gin.Default()
    err := router.SetTrustedProxies(nil)
    if err != nil {
        panic(err.Error())
    }

    router.LoadHTMLGlob("web/partials/*")

    router.Static("/styles", "web/styles")
    router.Static("/js", "web/js")

    router.GET("/", web.IndexHandler())
    router.GET("/report", web.ReportHandler())

    err = router.Run(":8080")
    if err != nil {
        panic(err)
    }
}