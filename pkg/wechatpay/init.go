package wechatpay

import (
	"context"
	"errors"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"log"
	"os"
	"strings"
)

var WechatPayClient *core.Client
var WechatPayJsAPI *jsapi.JsapiApiService
var WechatPayHandler *notify.Handler

func Init() error {
	var (
		mchID                      = os.Getenv("WC_ID")         // 商户号
		mchCertificateSerialNumber = os.Getenv("WC_SERIAL_NUM") // 商户证书序列号
		mchAPIv3Key                = os.Getenv("WC_API_V3_KEY") // 商户APIv3密钥
	)

	key := strings.Replace(os.Getenv("WC_CLIENT_KEY"), "\\n", "\n", -1)
	log.Println("WC_CLIENT_KEY:", key)
	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKey(key)
	if err != nil {
		return err
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	WechatPayClient, err = core.NewClient(ctx, opts...)
	if err != nil {
		return err
	}
	if WechatPayClient == nil {
		return errors.New("wechat pay init fail")
	}

	WechatPayJsAPI = &jsapi.JsapiApiService{Client: WechatPayClient}
	if WechatPayJsAPI == nil {
		return errors.New("wechat pay init fail")
	}

	certificateVisitor := downloader.MgrInstance().GetCertificateVisitor(mchID)
	// 3. 使用证书访问器初始化 `notify.Handler`
	WechatPayHandler, err = notify.NewRSANotifyHandler(mchAPIv3Key, verifiers.NewSHA256WithRSAVerifier(certificateVisitor))
	if err != nil {
		return err
	}
	if WechatPayHandler == nil {
		return errors.New("wechat pay init fail")
	}
	return nil
}
