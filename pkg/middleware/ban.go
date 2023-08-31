package middleware

import (
	"github.com/Zaprit/CrashReporter/pkg/config"
	"github.com/Zaprit/CrashReporter/pkg/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/netip"
)

func BanMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		addr, err := netip.ParseAddr(ctx.ClientIP())
		if err != nil {
			ctx.AbortWithStatus(500)
		}
		if db.IsBanned(addr) {
			ctx.HTML(http.StatusForbidden, "banned.gohtml", gin.H{"AppealURL": config.LoadedConfig.AppealURL})
			ctx.Abort()
		}
	}
}
