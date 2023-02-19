package api

import (
    "github.com/Zaprit/CrashReporter/pkg/lighthouse_client"
    "github.com/gin-gonic/gin"
)

type uriData struct {
    username string `uri:"username" binding:"required"`
}

func LighthouseUsersApiHandler() gin.HandlerFunc {
    return func(context *gin.Context) {
        var uridata uriData

        err := context.BindUri(&uridata)
        if err != nil {
            _ = context.Error(err)
        }

        data, status, err := lighthouse_client.GetUserRaw(uridata.username)
        if err != nil {
            context.String(status, err.Error())
        }
        context.String(status, string(data))
    }
}