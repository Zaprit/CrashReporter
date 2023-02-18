package web

import "github.com/gin-gonic/gin"

func SubmitReportHandler() gin.HandlerFunc {
    return func(context *gin.Context) {
        context.String(200, "Didn't bother to implement this yet :P")
    }
}