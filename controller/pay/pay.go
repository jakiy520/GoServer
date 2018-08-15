package pay

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"rggy/controller/common"

	"github.com/kataras/iris"
)

const (
	wxAppID = ""
	wxMchID = ""
	wxKey   = ""
)

// WXPayResp 响应信息
type WXPayResp struct {
	ReturnCode string
	ReturnMsg  string
	NonceStr   string
	PrepayID   string
}

// WxPay 微信支付
func WxPay(ctx iris.Context) {
	SendErrJSON := common.SendErrJSON

	info := make(map[string]interface{}, 0)

	fmt.Println("访问ip", ctx.RemoteAddr())
	ip := ctx.RemoteAddr()

	totalFee, _ := strconv.ParseFloat(ctx.Params().Get("total_fee"), 64) //单位 分
	openID := ctx.Params().Get("openId")                                 //"oKYr_0GkE-Izt9N9Wn43sapI9Pqw"
	body := "费用说明"
	//订单号
	orderNo := ctx.Params().Get("orderNo") //"wx"+utils.ToStr(time.Now().Unix()) + string(utils.Krand(4, 0))
	//随机数
	nonceStr := time.Now().Format("20060102150405") //+ string(utils.Krand(4, 0))
	var reqMap = make(map[string]interface{}, 0)
	reqMap["appid"] = wxAppID                                     //微信小程序appid
	reqMap["body"] = body                                         //商品描述
	reqMap["mch_id"] = wxMchID                                    //商户号
	reqMap["nonce_str"] = nonceStr                                //随机数
	reqMap["notify_url"] = "http://test.com.cn/weixinNotice.jspx" //通知地址
	reqMap["openid"] = openID                                     //商户唯一标识 openid
	reqMap["out_trade_no"] = orderNo                              //订单号
	reqMap["spbill_create_ip"] = ip                               //用户端ip   //订单生成的机器 IP
	reqMap["total_fee"] = totalFee * 100                          //订单总金额，单位为分
	reqMap["trade_type"] = "JSAPI"                                //trade_type=JSAPI时（即公众号支付），此参数必传，此参数为微信用户在商户对应appid下的唯一标识
	reqMap["sign"] = WxPayCalcSign(reqMap, wxKey)

	reqStr := Map2Xml(reqMap)
	fmt.Println("请求xml", reqStr)

	client := &http.Client{}

	// 调用支付统一下单API
	req, err := http.NewRequest("POST", "https://api.mch.weixin.qq.com/pay/unifiedorder", strings.NewReader(reqStr))
	if err != nil {
		// handle error
	}
	req.Header.Set("Content-Type", "text/xml;charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		SendErrJSON("获取预下单数据失败！", ctx)
		return
	}
	defer resp.Body.Close()

	body2, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		SendErrJSON("解析响应内容失败", ctx)
		return
	}
	fmt.Println("响应数据", string(body2))

	var resp1 WXPayResp
	err = xml.Unmarshal(body2, &resp1)
	if err != nil {
		panic(err)
	}

	// 返回预付单信息
	if strings.ToUpper(resp1.ReturnCode) == "SUCCESS" {
		fmt.Println("预支付申请成功")
		// 再次签名
		var resMap = make(map[string]interface{}, 0)
		resMap["appId"] = wxAppID
		resMap["nonceStr"] = resp1.NonceStr                            //商品描述
		resMap["package"] = "prepay_id=" + resp1.PrepayID              //商户号
		resMap["signType"] = "MD5"                                     //签名类型
		resMap["timeStamp"] = strconv.FormatInt(time.Now().Unix(), 10) //当前时间戳

		resMap["paySign"] = WxPayCalcSign(resMap, wxKey)
		// 返回5个支付参数及sign 用户进行确认支付

		fmt.Println("支付参数", resMap)
		// index.Console(resMap)
	} else {
		info["msg"] = "微信请求支付失败"
		// index.Console(info)
	}
}

// WxPayCalcSign 微信支付计算签名的函数
func WxPayCalcSign(mReq map[string]interface{}, key string) (sign string) {
	//STEP 1, 对key进行升序排序.
	sortedKeys := make([]string, 0)
	for k := range mReq {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	//STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sortedKeys {
		// logger.Printf("k=%v, v=%v\n", k, mReq[k])
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}

	//STEP3, 在键值对的最后加上key=API_KEY
	if key != "" {
		signStrings = signStrings + "key=" + key
	}

	fmt.Println("加密前-----", signStrings)
	//STEP4, 进行MD5签名并且将所有字符转为大写.
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings)) //
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))

	fmt.Println("加密后-----", upperSign)
	return upperSign
}

// Map2Xml 微信支付计算签名的函数
func Map2Xml(mReq map[string]interface{}) (xml string) {
	sb := bytes.Buffer{}
	sb.WriteString("<xml>")
	for k, v := range mReq {
		sb.WriteString("<" + k + ">" + v.(string) + "</" + k + ">")
	}
	sb.WriteString("</xml>")
	return sb.String()
}
