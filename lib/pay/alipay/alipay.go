package alipay

import (
	"fmt"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/xlog"
	"log"
	"go-admin/lib/common"
)

type AlipayMap struct {
	Alipay map[string]ClientConf
}
type ClientConf struct {
	AppId     string
	NotifyUrl string
	ReturnUrl string
	Charset   string
	SignType  string
	//key 路径
	AlipayPublicKey    string
	MerchantPrivateKey string
	AppPublicKey       string
	AppPrivateKey      string
	AlipayRootCert     string
	GatewayUrl         string
}

var conf AlipayMap

func New(app string) *alipay.Client {
	common.ParseConfig("alipay", &conf)

	c, ok := conf.Alipay[app]
	if !ok {
		log.Printf("alipay config is error")
		return nil
	}

	client, err := alipay.NewClient(c.AppId, c.MerchantPrivateKey, false)
	if err != nil {
		xlog.Error(err)
		return nil
	}
	//    注意：具体设置哪些参数，根据不同的方法而不同，此处列举出所有设置参数
	client.SetLocation(alipay.LocationShanghai). // 设置时区，不设置或出错均为默认服务器时间
							SetCharset(c.Charset).     // 设置字符编码，不设置默认 utf-8
							SetSignType(c.SignType).   // 设置签名类型，不设置默认 RSA2
							SetReturnUrl(c.ReturnUrl). // 设置返回URL
							SetNotifyUrl(c.NotifyUrl). // 设置异步通知URL
							SetAppAuthToken("")        // 设置第三方应用授权

	// 自动同步验签（只支持证书模式）
	// 传入 alipayCertPublicKey_RSA2.crt 内容
	client.AutoVerifySign([]byte("alipayCertPublicKey_RSA2 bytes"))

	// 公钥证书模式，需要传入证书，以下两种方式二选一
	// 证书路径
	err = client.SetCertSnByPath("appCertPublicKey.crt", "alipayRootCert.crt", "alipayCertPublicKey_RSA2.crt")
	fmt.Println(err)
	return client
}

func Test() {
	// 初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用私钥，支持PKCS1和PKCS8
	//    isProd：是否是正式环境
	privateKey := "privateKey"
	client, err := alipay.NewClient("2016091200494382", privateKey, false)
	if err != nil {
		xlog.Error(err)
		return
	}

	// 打开Debug开关，输出日志，默认关闭
	client.DebugSwitch = gopay.DebugOn

	// 设置支付宝请求 公共参数
	//    注意：具体设置哪些参数，根据不同的方法而不同，此处列举出所有设置参数
	client.SetLocation(alipay.LocationShanghai). // 设置时区，不设置或出错均为默认服务器时间
							SetCharset(alipay.UTF8).             // 设置字符编码，不设置默认 utf-8
							SetSignType(alipay.RSA2).            // 设置签名类型，不设置默认 RSA2
							SetReturnUrl("https://www.fmm.ink"). // 设置返回URL
							SetNotifyUrl("https://www.fmm.ink"). // 设置异步通知URL
							SetAppAuthToken("")                  // 设置第三方应用授权

	// 自动同步验签（只支持证书模式）
	// 传入 alipayCertPublicKey_RSA2.crt 内容
	client.AutoVerifySign([]byte("alipayCertPublicKey_RSA2 bytes"))

	// 公钥证书模式，需要传入证书，以下两种方式二选一
	// 证书路径
	err = client.SetCertSnByPath("appCertPublicKey.crt", "alipayRootCert.crt", "alipayCertPublicKey_RSA2.crt")
	fmt.Println(err)
	// 证书内容
	//err = client.SetCertSnByContent([]byte("appCertPublicKey bytes"), []byte("alipayRootCert bytes"), []byte("alipayCertPublicKey_RSA2 bytes"))

}
