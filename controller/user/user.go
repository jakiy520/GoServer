package user

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"net/http"
	"rggy/config"
	"rggy/controller/common"
	"rggy/model"
	"rggy/utils"
	"strings"
)

// WeAppLogin 微信小程序登录
func WeAppLogin(ctx iris.Context) {
	SendErrJSON := common.SendErrJSON
	code := ctx.FormValue("code")
	if code == "" {
		SendErrJSON("code不能为空", ctx)
		return
	}
	appID := config.WeAppConfig.AppID
	secret := config.WeAppConfig.Secret
	CodeToSessURL := config.WeAppConfig.CodeToSessURL
	CodeToSessURL = strings.Replace(CodeToSessURL, "{appid}", appID, -1)
	CodeToSessURL = strings.Replace(CodeToSessURL, "{secret}", secret, -1)
	CodeToSessURL = strings.Replace(CodeToSessURL, "{code}", code, -1)

	fmt.Println(CodeToSessURL)
	resp, err := http.Get(CodeToSessURL)
	if err != nil {
		fmt.Println(err.Error())
		SendErrJSON("error", ctx)
		return
	}
	fmt.Println(resp)

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		SendErrJSON("error", ctx)
		return
	}

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println(err.Error())
		SendErrJSON("error", ctx)
		return
	}

	if _, ok := data["session_key"]; !ok {
		fmt.Println("session_key 不存在")
		fmt.Println(data)
		SendErrJSON("error", ctx)
		return
	}

	var openID string
	var sessionKey string
	openID = data["openid"].(string)
	sessionKey = data["session_key"].(string)
	session := common.Sess.Start(ctx)
	session.Set("weAppOpenID", openID)
	session.Set("weAppSessionKey", sessionKey)

	resData := iris.Map{}
	resData[config.ServerConfig.SessionID] = session.ID()
	resData["openID"] = openID
	ctx.JSON(iris.Map{
		"errNo": model.ErrorCode.SUCCESS,
		"msg":   "success",
		"data":  resData,
	})
}

//	如果用户在数据库中不存在，新增user表
func newUser(weUser model.WeAppUser) (userid uint) {
	var user model.User

	if err := model.DB.First(&user, "open_id=?", weUser.OpenID).Error; err != nil {
		//	如果没有找到
		// fmt.Println(err)
		user.OpenID = weUser.OpenID
		user.Nickname = weUser.Nickname
		if weUser.Gender == 1 {
			user.Sex = true
		} else {
			user.Sex = false
		}
		model.DB.Create(&user)
	}
	userid = user.ID
	return
}

// SetWeAppUserInfo 设置小程序用户加密信息
func SetWeAppUserInfo(ctx iris.Context) {
	SendErrJSON := common.SendErrJSON
	type EncryptedUser struct {
		EncryptedData string `json:"encryptedData"`
		IV            string `json:"iv"`
	}
	var weAppUser EncryptedUser

	if ctx.ReadJSON(&weAppUser) != nil {
		SendErrJSON("参数错误", ctx)
		return
	}
	session := common.Sess.Start(ctx)
	sessionKey := session.GetString("weAppSessionKey")
	if sessionKey == "" {
		SendErrJSON("session error", ctx)
		return
	}

	userInfoStr, err := utils.DecodeWeAppUserInfo(weAppUser.EncryptedData, sessionKey, weAppUser.IV)
	if err != nil {
		fmt.Println(err.Error())
		SendErrJSON("error", ctx)
		return
	}

	var user model.WeAppUser
	if err := json.Unmarshal([]byte(userInfoStr), &user); err != nil {
		SendErrJSON("error", ctx)
		return
	}
	//	新增用户
	userid := newUser(user)

	session.Set("weAppUser", user)
	ctx.JSON(iris.Map{
		"errNo": model.ErrorCode.SUCCESS,
		"msg":   "success",
		"data":  iris.Map{"userid": userid},
	})
	return
}
