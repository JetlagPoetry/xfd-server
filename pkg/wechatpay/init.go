package wechatpay

import (
	"context"
	"errors"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"os"
)

var WechatPayClient *core.Client

func Init() error {
	var (
		mchID                      = os.Getenv("wx.mchID")                      // 商户号
		mchCertificateSerialNumber = os.Getenv("wx.mchCertificateSerialNumber") // 商户证书序列号
		mchAPIv3Key                = os.Getenv("wx.mchAPIv3Key")                // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
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
	return nil
}
