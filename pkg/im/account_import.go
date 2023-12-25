package im

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tencentyun/tls-sig-api-v2-golang/tencentyun"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"xfd-backend/pkg/utils"
	"xfd-backend/pkg/xerr"
)

func ImportAccount(userID string, phone string) (*Resp, error) {
	resp := &Resp{}
	imAppID, err := strconv.Atoi(os.Getenv("IM_APP_ID"))
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, err)
	}
	imSecret := os.Getenv("IM_APP_SECRET")
	imAdmin := os.Getenv("IM_ADMIN_USER")

	userSig, err := tencentyun.GenUserSig(imAppID, imSecret, imAdmin, 86400)
	if err != nil {
		return nil, err
	}
	fullURL := fmt.Sprintf("https://console.tim.qq.com/v4/im_open_login_svc/account_import?"+
		"sdkappid=%d&identifier=%s&usersig=%s&random=%d&contenttype=json", imAppID, imAdmin, userSig, rand.Uint32())
	body := map[string]string{
		"UserID":  userID,
		"Nick":    utils.GenUsername(phone),
		"FaceUrl": "https://xfd-t-1313159791.cos.ap-beijing.myqcloud.com/resources/common/aagj/WechatIMG1463.jpeg",
	}

	defer func() {
		log.Printf("[WxService] ImportAccount called, url=%s, resp=%v, err=%v\n", fullURL, utils.ToJson(resp), err)
	}()

	jsonBody, err := json.Marshal(body)
	if err != nil {
		fmt.Println("[Post] Marshal request failed, err = ", err)
		return nil, err
	}

	client := &http.Client{}
	httpReq, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Println("[Post] Create request failed, err = ", err)
		return nil, err
	}

	response, err := client.Do(httpReq)
	if err != nil {
		log.Println("[Post] Make request failed, err = ", err)
		return nil, err
	}

	// 只有200才返回成功
	if response.StatusCode != 200 {
		log.Println("[Post] response status code != 200")
		return nil, errors.New("http call failed")
	}

	defer response.Body.Close()
	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("[Post] Read response failed, err = ", err)
		return nil, err
	}

	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		log.Println("[Post] Read response failed, err = ", err)
		return nil, err
	}
	if resp.ErrorCode != 0 {
		return nil, errors.New("API call failed")
	}
	return resp, nil

}

type Resp struct {
	ActionStatus string `json:"ActionStatus"`
	ErrorInfo    string `json:"ErrorInfo"`
	ErrorCode    int    `json:"ErrorCode"`
}
