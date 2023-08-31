package api

import (
	"github.com/Zaprit/CrashReporter/pkg/db"
	"github.com/gin-gonic/gin"
	"net/netip"
	"time"
)

type BanModel struct {
	IP       string
	Duration int
	Reason   string
}

type BanResponse struct {
	Success bool
	Message string
}

func BanHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		model := BanModel{}
		err := ctx.BindJSON(&model)
		if err != nil {
			ctx.Error(err)
			ctx.JSON(400, &BanResponse{
				Success: false,
				Message: "Invalid data",
			})
		}

		addr, err := netip.ParseAddr(model.IP)
		if err != nil {
			ctx.Error(err)
			ctx.JSON(400, &BanResponse{
				Success: false,
				Message: "Invalid IP",
			})
		}

		banDays := time.Duration(model.Duration * 24)

		banTime := time.Now().Add(time.Hour * banDays)

		db.BanIP(addr, model.Reason, banTime)
	}
}
