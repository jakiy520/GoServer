package product

import (
	"encoding/json"
	"fmt"
	"rggy/config"
	"rggy/controller/common"
	"rggy/controller/kanjia"
	"rggy/model"

	"github.com/kataras/iris"
	// "time"
)

//GetKanjiaPro 获取当前的活动商品
func GetKanjiaPro(ctx iris.Context) {
	SendErrJSON := common.SendErrJSON
	// reqStartTime := time.Now()

	//	获取商品基础信息
	// id, err := ctx.ParamInt("id")
	productID := uint(config.KanjiaConfig.ProductID)
	// var productId uint = 1
	var product model.Product

	//	获取用户id
	userID, err := ctx.Params().GetInt("userID")
	if err != nil {
		SendErrJSON("用户id参数有误", ctx)
		return
	}

	//	获取砍价数据
	kanjiaID, err := ctx.Params().GetInt("kanjiaID")
	if err != nil {
		SendErrJSON("砍价id参数有误", ctx)
		return
	}
	var modelKanjia model.Kanjia
	var kanjiaMoney float64
	//	如果没有传入 kanjiaID ， 根据用户id和商品id 获取该用户有没有参与该商品的砍价
	if userID > 0 && kanjiaID == 0 {
		if err := model.DB.First(&modelKanjia, "user_id=? and product_id=?", userID, productID).Error; err == nil {
		}
	}
	if kanjiaID > 0 {
		if err := model.DB.First(&modelKanjia, "id=?", kanjiaID).Error; err != nil {
			SendErrJSON("不存在该砍价编号！", ctx)
			return
		}
	}
	if modelKanjia.ID > 0 {
		productID = modelKanjia.ProductID
		kanjiaMoney = kanjia.GetKanjiaMoney(uint(kanjiaID))
	}

	//	初始化商品信息
	if model.DB.First(&product, productID).Error != nil {
		SendErrJSON("错误的商品id", ctx)
		return
	}

	if model.DB.First(&product.Image, product.ImageID).Error != nil {
		product.Image = model.Image{}
	}

	var imagesSQL []uint
	if err := json.Unmarshal([]byte(product.ImageIDs), &imagesSQL); err == nil {
		var images []model.Image
		if model.DB.Where("id in (?)", imagesSQL).Find(&images).Error != nil {
			product.Images = nil
		} else {
			product.Images = images
		}
	} else {
		product.Images = nil
	}

	if err := model.DB.Model(&product).Related(&product.Categories, "categories").Error; err != nil {
		fmt.Println(err.Error())
		SendErrJSON("error", ctx)
		return
	}

	if product.HasProperty {
		if err := model.DB.Model(&product).Related(&product.Properties).Error; err != nil {
			fmt.Println(err.Error())
			SendErrJSON("error", ctx)
			return
		}

		for i := 0; i < len(product.Properties); i++ {
			property := product.Properties[i]
			if err := model.DB.Model(&property).Related(&property.PropertyValues).Error; err != nil {
				fmt.Println(err.Error())
				SendErrJSON("error", ctx)
				return
			}
			product.Properties[i] = property
		}

		if err := model.DB.Model(&product).Related(&product.Inventories).Error; err != nil {
			fmt.Println(err.Error())
			SendErrJSON("error", ctx)
			return
		}

		for i := 0; i < len(product.Inventories); i++ {
			inventory := product.Inventories[i]
			if err := model.DB.Model(&inventory).Related(&inventory.PropertyValues, "property_values").Error; err != nil {
				fmt.Println(err.Error())
				SendErrJSON("error", ctx)
				return
			}
			product.Inventories[i] = inventory
		}
	}
	// fmt.Println(product)
	// fmt.Println("duration: ", time.Now().Sub(reqStartTime).Seconds())

	ctx.JSON(iris.Map{
		"errNo": model.ErrorCode.SUCCESS,
		"msg":   "success",
		"data": iris.Map{
			"product":     product,
			"kanjiaInfo":  modelKanjia,
			"kanjiaMoney": fmt.Sprintf("%.2f", kanjiaMoney),
		},
	})
}
