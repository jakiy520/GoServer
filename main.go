package main

import (
	"fmt"
	"os"
	"rggy/config"
	"rggy/controller/common"
	"rggy/model"
	"rggy/route"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

func init() {
	db, err := gorm.Open(config.DBConfig.Dialect, config.DBConfig.URL)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	if config.DBConfig.SQLLog {
		db.LogMode(true)
	}

	db.DB().SetMaxIdleConns(config.DBConfig.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.DBConfig.MaxOpenConns)

	model.DB = db

	common.Sess = sessions.New(sessions.Config{
		Cookie:  config.ServerConfig.SessionID,
		Expires: time.Minute * 20,
	})

	initDBTables()
}

//	初始化数据库表结构
func initDBTables() {
	// Migrate the schema11111
	model.DB.AutoMigrate(&model.User{})
	model.DB.AutoMigrate(&model.Product{})
	model.DB.AutoMigrate(&model.Image{})
	model.DB.AutoMigrate(&model.Inventory{})
	model.DB.AutoMigrate(&model.Category{})
	model.DB.AutoMigrate(&model.Property{})
	model.DB.AutoMigrate(&model.PropertyValue{})
	model.DB.AutoMigrate(&model.Kanjia{})
	model.DB.AutoMigrate(&model.KanjiaRecord{})
}

func main() {

	app := iris.New()

	route.Route(app)
	app.StaticWeb("/images", "./public/images")

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.Writef("404 not found here1111")
	})
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.HTML("Message: <b>" + ctx.Values().GetString("message") + "</b>")
	})
	app.Run(iris.Addr(":8081"), iris.WithConfiguration(iris.Configuration{ // default configuration:
		Charset: "UTF-8",
	}))

}
