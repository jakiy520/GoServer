package main

import (
	"github.com/kataras/iris"
	"rggyServer/route"
)

func main() {
	app := iris.New()

	route.Route(app)

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.Writef("404 not found here1111")
	})
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.HTML("Message: <b>" + ctx.Values().GetString("message") + "</b>")
	})
	app.Run(iris.Addr(":8081"))
}
