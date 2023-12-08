package wechatpay

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	"xfd-backend/pkg/consts"
	"xfd-backend/pkg/types"
	"xfd-backend/pkg/utils"
	"xfd-backend/pkg/xerr"
)

func GetWxOpenID(ctx context.Context, code string) (*types.WxOpenIDResp, xerr.XErr) {
	var (
		result *types.WxOpenIDResp
		err    error
	)
	url := "https://api.weixin.qq.com/sns/jscode2session"
	appID := os.Getenv("WECHAT_APP_ID")
	secret := os.Getenv("WECHAT_APP_SECRET")
	grantType := "authorization_code"

	fullURL := fmt.Sprintf("%s?appid=%s&secret=%s&grant_type=%s&js_code=%s", url, appID, secret, grantType, code)

	defer func() {
		log.Printf("[WxService] GetWxOpenID  called, url=%s, resp=%v, err=%v\n", fullURL, utils.ToJson(result), err)
	}()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Get(fullURL)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorCallApi, err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorCallApi, err)
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorCallApi, err)
	}

	if len(result.OpenID) == 0 {
		log.Println("[GetWxOpenID] get openID failed, openID is empty")
		return nil, xerr.WithCode(xerr.ErrorCallApi, errors.New("openid empty"))
	}
	return result, nil
}

func CreateOrder(ctx context.Context, orderSn, desc, openID string, price int64) (*jsapi.PrepayWithRequestPaymentResponse, xerr.XErr) {
	var (
		req  jsapi.PrepayRequest
		resp *jsapi.PrepayWithRequestPaymentResponse
		err  error
	)
	defer func() {
		log.Printf("[WxService] CreateOrder called, req=%s, resp=%v, err=%v\n", utils.ToJson(req), utils.ToJson(resp), err)
	}()

	// 测试环境，强制用1分钱
	if os.Getenv("WECHAT_TEST_OPEN") == "true" {
		price = 1
	}

	// 得到prepay_id，以及调起支付所需的参数和签名
	req = jsapi.PrepayRequest{
		Appid:       core.String(os.Getenv("APP_ID")),
		Mchid:       core.String(os.Getenv("WC_ID")),
		Description: core.String(desc),
		OutTradeNo:  core.String(orderSn),
		//Attach:      core.String("自定义数据说明"),
		TimeExpire: core.Time(time.Now().Add(time.Minute * consts.WECHAT_PAY_EXPIRE_MINUTE)),
		NotifyUrl:  core.String(os.Getenv("DOMAIN_NAME") + "/api/v1/order/paymentConfirm"),
		Amount: &jsapi.Amount{
			Total: core.Int64(price),
		},
		Payer: &jsapi.Payer{
			Openid: core.String(openID),
		},
	}
	resp, _, err = WechatPayJsAPI.PrepayWithRequestPayment(ctx, req)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorCallApi, err)
	}
	return resp, nil
}

func CancelOrder(ctx context.Context, orderSn string) xerr.XErr {
	var (
		req  jsapi.CloseOrderRequest
		resp *payments.Transaction
		err  error
	)

	defer func() {
		log.Printf("[WxService] CancelOrder called, req=%s, resp=%v, err=%v\n", utils.ToJson(req), utils.ToJson(resp), err)
	}()

	req = jsapi.CloseOrderRequest{
		OutTradeNo: core.String(orderSn),
		Mchid:      core.String(os.Getenv("WC_ID")),
	}
	_, err = WechatPayJsAPI.CloseOrder(ctx, req)

	return nil
}

func LookupOrder(ctx context.Context, orderSn string) (*payments.Transaction, xerr.XErr) {
	var (
		req  jsapi.QueryOrderByOutTradeNoRequest
		resp *payments.Transaction
		err  error
	)

	defer func() {
		log.Printf("[WxService] LookupOrder called, req=%s, resp=%v, err=%v\n", utils.ToJson(req), utils.ToJson(resp), err)
	}()

	req = jsapi.QueryOrderByOutTradeNoRequest{
		OutTradeNo: core.String(orderSn),
		Mchid:      core.String(os.Getenv("WC_ID")),
	}
	resp, _, err = WechatPayJsAPI.QueryOrderByOutTradeNo(ctx, req)

	return resp, nil
}
