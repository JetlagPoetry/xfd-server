package wechatpay

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

	resp, err := http.Get(fullURL)
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

func CreateOrder(ctx context.Context, orderSn, desc, openID string, price int64) (*types.WxOrderResponse, xerr.XErr) {
	var (
		req  jsapi.PrepayRequest
		resp *jsapi.PrepayWithRequestPaymentResponse
		err  error
	)
	defer func() {
		log.Printf("[WxService] CreateOrder called, url=%s, resp=%v, err=%v\n", utils.ToJson(req), utils.ToJson(resp), err)
	}()
	svc := jsapi.JsapiApiService{Client: WechatPayClient}
	// 得到prepay_id，以及调起支付所需的参数和签名
	req = jsapi.PrepayRequest{
		Appid:       core.String(os.Getenv("WC_APP_ID")),
		Mchid:       core.String(os.Getenv("WC_MCH_ID")),
		Description: core.String(desc),
		OutTradeNo:  core.String(orderSn),
		//Attach:      core.String("自定义数据说明"),
		NotifyUrl: core.String("https://www.weixin.qq.com/wxpay/pay.php"),
		Amount: &jsapi.Amount{
			Total: core.Int64(price),
		},
		Payer: &jsapi.Payer{
			Openid: core.String(openID),
		},
	}
	resp, _, err = svc.PrepayWithRequestPayment(ctx, req)
	if err == nil {
		return nil, xerr.WithCode(xerr.ErrorCallApi, err)
	}
	return &types.WxOrderResponse{}, nil
}
