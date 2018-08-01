package route

import (
	"rggy/controller/kanjia"

	"github.com/kataras/iris"

	// "rggy/controller/pay"
	"rggy/controller/product"
	"rggy/controller/user"
)

//Route 1
func Route(app *iris.Application) {

	router := app.Party("/api")
	{
		router.Get("/", func(ctx iris.Context) {
			ctx.Text("hello ")
		})
		router.Get("/weAppLogin", user.WeAppLogin)
		router.Post("/setWeAppUser", user.SetWeAppUserInfo)
		router.Get("/getKanjiaPro/:kanjiaID/:userID", product.GetKanjiaPro)
		router.Post("/JoinKanjia", kanjia.JoinKanjia)
		router.Post("/Bangtakan", kanjia.Bangtakan)
		router.Get("/GetKanjiaRecords/:kanjiaID", kanjia.GetKanjiaRecords)
	}
}
