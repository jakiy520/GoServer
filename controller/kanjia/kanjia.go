package kanjia

import (
	"fmt"
	"math/rand"
	"rggy/controller/common"
	"rggy/model"
	"rggy/utils"
	"strconv"

	"github.com/kataras/iris"
)

//JoinKanjia 参与砍价活动
func JoinKanjia(ctx iris.Context) {
	SendErrJSON := common.SendErrJSON
	type ModelJoinKanjia struct {
		UserID        uint   `json:"userID"`
		UserNickName  string `json:"userNickName"`
		UserAvatarURL string `json:"userAvatarUrl"`
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
		kanjia.UserAvatarURL = modelJoin.UserAvatarURL
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

//Bangtakan 帮他砍价
func Bangtakan(ctx iris.Context) {
	SendErrJSON := common.SendErrJSON
	type ModelBangtakan struct {
		KanjiaID      uint   `json:"kanjiaID"`
		UserID        uint   `json:"userID"`
		UserNickName  string `json:"userNickName"`
		UserAvatarURL string `json:"userAvatarUrl"`
		ProductID     uint   `json:"productID"`
	}
	var modelBangtakan ModelBangtakan

	if ctx.ReadJSON(&modelBangtakan) != nil {
		SendErrJSON("参数错误", ctx)
		return
	}

	//	获取砍价商品基本信息
	var modelProduct model.Product
	if err := model.DB.First(&modelProduct, "id=?", modelBangtakan.ProductID).Error; err != nil {
		SendErrJSON("该砍价商品不存在", ctx)
		return
	}

	//	生成砍价记录
	var modelKanjiaRecord model.KanjiaRecord
	//	判断该用户是否已经砍过了
	if err := model.DB.First(&modelKanjiaRecord, "kanjia_id=? and user_id=?", modelBangtakan.KanjiaID, modelBangtakan.UserID).Error; err != nil {
		modelKanjiaRecord.KanjiaID = modelBangtakan.KanjiaID
		modelKanjiaRecord.UserID = modelBangtakan.UserID
		modelKanjiaRecord.UserNickName = modelBangtakan.UserNickName
		modelKanjiaRecord.UserAvatarURL = modelBangtakan.UserAvatarURL
		modelKanjiaRecord.ProductID = modelBangtakan.ProductID
		//	获取一个随机的砍价金额，最大值不能超过数据库设置的砍价单次最大值
		_kanjiamoney := (modelProduct.KanjiaMaxMoneyOne * rand.Float64())
		_kanjiamoney, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", _kanjiamoney), 64)
		modelKanjiaRecord.KanjiaPrice = _kanjiamoney
	} else {
		SendErrJSON("该用户已经砍过了", ctx)
		return
	}

	//	砍价总金额
	allKanjiaMoney := GetKanjiaMoney(modelBangtakan.KanjiaID)
	allKanjiaMoney += modelKanjiaRecord.KanjiaPrice

	//	如果砍价金额大于最低值，则不允许再砍了
	if allKanjiaMoney > modelProduct.KanjiaMaxMoneyAll {
		SendErrJSON("不能再砍了", ctx)
		return
	}

	if modelProduct.Price-allKanjiaMoney <= 0 {
		SendErrJSON("不能再砍了", ctx)
	}

	model.DB.Create(&modelKanjiaRecord)

	ctx.JSON(iris.Map{
		"errNo": model.ErrorCode.SUCCESS,
		"msg":   "success",
		"data":  iris.Map{"AllKanjiaMoney": fmt.Sprintf("%.2f", allKanjiaMoney)},
	})
	return
}

//GetKanjiaMoney 获取已砍价价格
func GetKanjiaMoney(kanjiaID uint) (allKanjiaMoney float64) {
	var listKanjiaRecord []model.KanjiaRecord
	if err := model.DB.Where("kanjia_id=?", kanjiaID).Find(&listKanjiaRecord).Error; err == nil {
		//	如果查询到砍价记录，则判断砍价金额
		for i := 0; i < len(listKanjiaRecord); i++ {
			kanjiaRecord := listKanjiaRecord[i]
			allKanjiaMoney += kanjiaRecord.KanjiaPrice
		}

	}
	return
}

//GetKanjiaRecords 获取砍价记录
func GetKanjiaRecords(ctx iris.Context) {
	SendErrJSON := common.SendErrJSON

	kanjiaID, err := ctx.Params().GetInt("kanjiaID")
	if err != nil {
		SendErrJSON("砍价id参数有误", ctx)
		return
	}

	//	获取砍价商品基本信息
	var listKanjiaRecords []model.KanjiaRecord
	if err := model.DB.Where("kanjia_id=?", kanjiaID).Find(&listKanjiaRecords).Error; err != nil {

	}
	ctx.JSON(iris.Map{
		"errNo": model.ErrorCode.SUCCESS,
		"msg":   "success",
		"data":  iris.Map{"KanjiaRecords": listKanjiaRecords},
	})
	return
}
