package route

import (
	"github.com/kataras/iris"
)

func Route(app *iris.Application) {
	adminRoutes := app.Party("/admin")
	adminRoutes.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome admin. </h1>")
	})

	// Method:   GET
	// Resource: http://localhost:8080ttt
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	// Method:   GET
	// Resource: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})
}
