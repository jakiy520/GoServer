// package pay

// import (
// 	"crypto/md5"
// 	"encoding/hex"
// 	"fmt"
// 	"sort"
// 	"strings"
// )

// //首先定义一个UnifyOrderReq用于填入我们要传入的参数。
// type UnifyOrderReq struct {
// 	Appid            string `xml:"appid"`
// 	Body             string `xml:"body"`
// 	Mch_id           string `xml:"mch_id"`
// 	Nonce_str        string `xml:"nonce_str"`
// 	Notify_url       string `xml:"notify_url"`
// 	Trade_type       string `xml:"trade_type"`
// 	Spbill_create_ip string `xml:"spbill_create_ip"`
// 	Total_fee        int    `xml:"total_fee"`
// 	Out_trade_no     string `xml:"out_trade_no"`
// 	Sign             string `xml:"sign"`
// }

// //微信支付计算签名的函数
// func wxpayCalcSign(mReq map[string]interface{}, key string) (sign string) {
// 	fmt.Println("微信支付签名计算, API KEY:", key)
// 	//STEP 1, 对key进行升序排序.
// 	sorted_keys := make([]string, 0)
// 	for k, _ := range mReq {
// 		sorted_keys = append(sorted_keys, k)
// 	}

// 	sort.Strings(sorted_keys)

// 	//STEP2, 对key=value的键值对用&连接起来，略过空值
// 	var signStrings string
// 	for _, k := range sorted_keys {
// 		fmt.Printf("k=%v, v=%v\n", k, mReq[k])
// 		value := fmt.Sprintf("%v", mReq[k])
// 		if value != "" {
// 			signStrings = signStrings + k + "=" + value + "&"
// 		}
// 	}

// 	//STEP3, 在键值对的最后加上key=API_KEY
// 	if key != "" {
// 		signStrings = signStrings + "key=" + key
// 	}

// 	//STEP4, 进行MD5签名并且将所有字符转为大写.
// 	md5Ctx := md5.New()
// 	md5Ctx.Write([]byte(signStrings))
// 	cipherStr := md5Ctx.Sum(nil)
// 	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
// 	return upperSign
// }

// //	执行统一下单
// func DoPay(ctx iris.Context) {
// 	ip := ctx.RemoteAddr()

// 	//请求UnifiedOrder的代码
// 	var yourReq UnifyOrderReq
// 	yourReq.Appid = "app_id" //微信开放平台我们创建出来的app的app id
// 	yourReq.Body = "商品名"
// 	yourReq.Mch_id = "商户编号"
// 	yourReq.Nonce_str = "your nonce"
// 	yourReq.Notify_url = "www.yourserver.com/wxpayNotify"
// 	yourReq.Trade_type = "APP"
// 	yourReq.Spbill_create_ip = "xxx.xxx.xxx.xxx"
// 	yourReq.Total_fee = 10 //单位是分，这里是1毛钱
// 	yourReq.Out_trade_no = "后台系统单号"

// 	var m map[string]interface{}
// 	m = make(map[string]interface{}, 0)
// 	m["appid"] = yourReq.Appid
// 	m["body"] = yourReq.Body
// 	m["mch_id"] = yourReq.Mch_id
// 	m["notify_url"] = yourReq.Notify_url
// 	m["trade_type"] = yourReq.Trade_type
// 	m["spbill_create_ip"] = yourReq.Spbill_create_ip
// 	m["total_fee"] = yourReq.Total_fee
// 	m["out_trade_no"] = yourReq.Out_trade_no
// 	m["nonce_str"] = yourReq.Nonce_str
// 	yourReq.Sign = wxpayCalcSign(m, "wxpay_api_key") //这个是计算wxpay签名的函数上面已贴出

// 	bytes_req, err := xml.Marshal(yourReq)
// 	if err != nil {
// 		fmt.Println("以xml形式编码发送错误, 原因:", err)
// 		return
// 	}

// 	str_req := string(bytes_req)
// 	//wxpay的unifiedorder接口需要http body中xmldoc的根节点是<xml></xml>这种，所以这里需要replace一下
// 	str_req = strings.Replace(str_req, "UnifyOrderReq", "xml", -1)
// 	bytes_req = []byte(str_req)

// 	//发送unified order请求.
// 	req, err := http.NewRequest("POST", unify_order_req, bytes.NewReader(bytes_req))
// 	if err != nil {
// 		fmt.Println("New Http Request发生错误，原因:", err)
// 		return
// 	}
// 	req.Header.Set("Accept", "application/xml")
// 	//这里的http header的设置是必须设置的.
// 	req.Header.Set("Content-Type", "application/xml;charset=utf-8")

// 	c := http.Client{}
// 	resp, _err := c.Do(req)
// 	if _err != nil {
// 		fmt.Println("请求微信支付统一下单接口发送错误, 原因:", _err)
// 		return
// 	}

// 	//到这里统一下单接口就已经执行完成了
// }
