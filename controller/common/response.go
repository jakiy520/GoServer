package common

import (
	"github.com/kataras/iris"
	"rggy/model"
)

// SendErrJSON 有错误发生时，发送错误JSON
func SendErrJSON(msg string, ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK) // default is 200 == iris.StatusOK
	ctx.JSON(iris.Map{
		"errNo": model.ErrorCode.ERROR,
		"msg":   msg,
		"data":  iris.Map{},
	})
}
