package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"os"
	"rggy/config"
	"rggy/controller/common"
	"rggy/model"
	"rggy/route"
	"time"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

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
}

func main() {
	// Migrate the schema11111
	model.DB.AutoMigrate(&Product{})

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
