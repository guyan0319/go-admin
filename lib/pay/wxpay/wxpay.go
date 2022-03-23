package alipay

import (
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xlog"
	"github.com/go-pay/gopay/wechat/v3"
	"log"
	"web-demo/lib/common"
)

type WxpayMap struct {
	Wxpay map[string]ClientConf
}
type ClientConf struct {
	AppId      string
	Mchid      string
	SerialNo   string
	APIv3Key   string
	PrivateKey string
}

var conf WxpayMap

func New(app string) *wechat.ClientV3 {
	common.ParseConfig("wxpay", &conf)
	c, ok := conf.Wxpay[app]
	if !ok {
		log.Printf("wxpay config is error")
		return nil
	}

	// NewClientV3 初始化微信客户端 v3
	//	mchid：商户ID 或者服务商模式的 sp_mchid
	// 	serialNo：商户证书的证书序列号
	//	apiV3Key：apiV3Key，商户平台获取
	//	privateKey：私钥 apiclient_key.pem 读取后的内容
	client, err := wechat.NewClientV3(c.Mchid, c.SerialNo, c.APIv3Key, c.PrivateKey)
	if err != nil {
		xlog.Error(err)
		return nil
	}

	// 启用自动同步返回验签，并定时更新微信平台API证书
	err = client.AutoVerifySign()
	if err != nil {
		xlog.Error(err)
		return nil
	}

	// 打开Debug开关，输出日志，默认是关闭的
	client.DebugSwitch = gopay.DebugOn
	return client
}
