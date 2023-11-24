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

	log.Println("WC_CLIENT_KEY:", os.Getenv("WC_CLIENT_KEY"))
	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKey("-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQC/Vt3txH4a2FAt\nrkOXvF4FXw36Je/Oelk0DzPSrjcZm3Zfh66rR7vkyQ+iBnXsh7Pgnl+CSvEeuCHJ\nlmhBK+Sn5SMXzHlfNGGDQPDOH/TQCjkJHLPNoZI5yRw/9SRBKi+l4hEfoE06H5Eg\nPj94vzrEQqPFRpdvXT/R+SL+JPjrdSPIP/3AJUKicGto6vaS0rqOoIt5uu9WhB5W\nLQY8eAVh5/vOs9wmYcxqL+FE62O04h05oQLgi1DOiozWh4ac9y7b53ohkcLpCSxm\nBD8fFyHQ6P3LdFwiRCHSv2GruvnRHuWdyX4Z+66Yyiw+xfYPRUPzlNbkit3yaG9r\nCzo0/AXhAgMBAAECggEAT5NZCSHHMHL2JYlsIxb8I0+9CGPur/jCUJxh+STrVYMK\nYACV7DYRNr1H+yKkuJJXhPtYzV9NmHEL3ELpaPFQT2NEjQlFEQs/4s6HV7KCpmMa\nBWgif0rK/a1eSlsxkMhyuJKkfHq2o42tVbONwjIFxsnQQqNxJ6eoezMaEohd3c5J\nLNQYbFoM9bn9WI44v0YdShPOjqyGzgEmMVZYl0SfZMyVg/hAh9CFoBEk/t1JBw6z\ntJ7WinalVSkzVoDZqX7+CdIAEF1yZehMwZTGKhIrUVthvmjD7Aj/xoTEFg429B5q\n5Sf1O8qHOqYhh8irvEerWlhJD6Ozek765r5ZL4dhYQKBgQDwwkszGcAjcKH0aBqc\nNwpyYmZhKy4K+dP+tVlT9zWb7eUO/TjfZrhYk9uWXxYUBdiLOYdmdAGh0C/o9DSp\nmXhqJ2dOaDsd4YdVI6TUXJ6r3Jhn2RKX6s+UQ2gmCt6PZS1tIWTiZ57Yw6lflfeP\nXS8ongh7GC79ASSdoPDlbWR3cwKBgQDLc692tw1/OY5EVdaTyd9OjszBnUR7N96+\n5oOBJtqoLmUcTRsRL5waCkMc/M3y2IRSefPJRR21y6YTUZDaD6lgdiJ0SVQ21+sp\nSzmMlU8CIaLVkrSy0HfSxJJrFwigMxZflMlZA0o/WtNxFYMnTn6hEDxNmP5ze6YN\ni/No3KAwWwKBgQCEu2lAA+tU3oy7NQWup/2fcDZnTan1rSQ+IXbc/sZUcAQ/jkVj\njLsHQoTmUwfWBB8NKqtGRaB8uE0hjjjWY+DqDU08AeTNpX+55YvC9EkaEOUJI7jH\nflHwuHTbvFRVE//GUYnP9Daz5LMlGoXASSxtpSqhyoRlQdMRACCvK+8elwKBgQC8\nVJhmMkkCLkstorABGmvmEnYj++q6jRtaZ5Pv6AGckWXzbsTRgdrkl/9MRBB47kh+\n+HjdJWe9M4jGdUi7Mqg9rN1z+7VDF1iOXx4krsn2VSCgxy8SH7vrlR9clnPbp67c\nR6SjC1KdlvwHwDwqFJVPjvHjeu6ABEPQYm6t8R0v5QKBgQC3vGevGO/IqVizld/s\nMIepI8GJATMQjIGCBJ6/JPiIN/9A0zCmj8x8CKinnstnh77q19XALVT1gHoIxGQE\nEjHFs4gnDocfW8nbZmIpklVzu5bxKtjC8egGmL3hy26SA0/y/XNKkzcBPQ4VoKtH\n9zKnk+pelGP4JoNHcWkt2G2yeg==\n-----END PRIVATE KEY-----\n")
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
