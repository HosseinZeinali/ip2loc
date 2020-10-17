package http

import (
	"github.com/HosseinZeinali/ip2loc/app"
)

type Api struct {
	Router *Router
	App    *app.App
}

func NewApi(router *Router, app *app.App) *Api {
	api := new(Api)
	api.Router = router
	api.App = app

	return api
}

func (a *Api) InitRoutes() {
	ctx := a.App.NewContext()
	a.Router.GET("/api/v1/ip/:ip", a.IpHandler(ctx))
}
