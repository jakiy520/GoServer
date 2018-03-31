package kanjia

import (
	"github.com/kataras/iris"
	"rggy/controller/common"
	"rggy/model"
	"rggy/utils"
)

// 参与砍价活动
func JoinKanjia(ctx iris.Context) {
	SendErrJSON := common.SendErrJSON
	type ModelJoinKanjia struct {
		UserID        uint   `json:"userID"`
		UserNickName  string `json:"userNickName"`
		UserAvatarUrl string `json:"userAvatarUrl"`
		ProductID     uint   `json:"productID"`
	}
	var modelJoin ModelJoinKanjia

	if ctx.ReadJSON(&modelJoin) != nil {
		SendErrJSON("参数错误", ctx)
		return
	}
	var kanjia model.Kanjia

	if err := model.DB.First(&kanjia, "user_id=? and product_id=?", modelJoin.UserID, modelJoin.ProductID).Error; err != nil {
		kanjia.UserID = modelJoin.UserID
		kanjia.ProductID = modelJoin.ProductID
		kanjia.UserNickName = modelJoin.UserNickName
		kanjia.UserAvatarUrl = modelJoin.UserAvatarUrl
		kanjia.ValidCode = utils.GetRandValidCode(6)
		model.DB.Create(&kanjia)
	}

	ctx.JSON(iris.Map{
		"errNo": model.ErrorCode.SUCCESS,
		"msg":   "success",
		"data":  iris.Map{"kanjiaID": kanjia.ID},
	})
	return
}
