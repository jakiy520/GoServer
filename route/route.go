package route

import (
	"github.com/kataras/iris"
	"rggy/controller/product"
	"rggy/controller/user"
)

func Route(app *iris.Application) {

	router := app.Party("/api")
	{
		router.Get("/", func(ctx iris.Context) {
			ctx.Text("hello ")
		})
		router.Get("/weAppLogin", user.WeAppLogin)
		router.Post("/setWeAppUser", user.SetWeAppUserInfo)
		router.Get("/getKanjiaPro", product.GetKanjiaPro)
	}
}
