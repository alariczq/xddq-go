package app

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/pprof"
)

func (a *App) initServer() {
	pprof.Register(a.server, "/pprof")

	a.server.POST("/stop", func(c context.Context, ctx *app.RequestContext) {
		a.stopAssists()
	})

	a.server.POST("/start", func(c context.Context, ctx *app.RequestContext) {
		a.startAssists()
	})

	a.server.GET("/activity", func(c context.Context, ctx *app.RequestContext) {
		for _, a := range a.assists {
			ctx.JSON(http.StatusOK, a.Activity().ListActivity())
		}
	})
	a.server.GET("/activity/detail", func(c context.Context, ctx *app.RequestContext) {
		for _, a := range a.assists {
			ctx.JSON(http.StatusOK, a.Activity().ListActivityDetail())
		}
	})

	a.server.GET("/yard", func(c context.Context, ctx *app.RequestContext) {
		for _, a := range a.assists {
			yard := a.Yard()
			ctx.JSON(http.StatusOK, yard.ListMerchantInfo())
			return
		}
	})
}
