package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris"
	"os"
	"rggyServer/config"
	"rggyServer/model"
	"rggyServer/route"
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
}

func main() {
	// Migrate the schema1212
	model.DB.AutoMigrate(&Product{})

	// Create
	model.DB.Create(&Product{Code: "L1212", Price: 1000})
	// Read
	var product Product
	model.DB.First(&product, 1)                   // find product with id 1
	model.DB.First(&product, "code = ?", "L1212") // find product with code l1212

	model.DB.Delete(&product)

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
