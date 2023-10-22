package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"xfd-backend/pkg/consts"
	"xfd-backend/pkg/xerr"
)

// Response 基础序列化器
type Response struct {
	LogID  string      `json:"log_id"`
	Status xerr.XCode  `json:"status"`
	ErrMsg string      `json:"err_msg"`
	Data   interface{} `json:"data"`
}

// TrackedErrorResponse 有追踪信息的错误反应
type TrackedErrorResponse struct {
	Response
	LogID string `json:"log_id"`
}

// RespSuccess 带data成功返回
func RespSuccess(ctx *gin.Context, data interface{}) *Response {
	logID := GetLogIDFromCtx(ctx)
	status := xerr.SUCCESS

	if data == nil {
		data = "操作成功"
	}

	r := &Response{
		Status: status,
		Data:   data,
		ErrMsg: GetMsg(status),
		LogID:  logID,
	}
	ctx.Set(consts.CONTEXT_HEADER_RESP_BODY, r.Data)
	ctx.Set(consts.CONTEXT_HEADER_RESP_MSG, r.ErrMsg)
	ctx.Set(consts.CONTEXT_HEADER_RESP_HTTP_CODE, http.StatusOK)
	ctx.Set(consts.CONTEXT_HEADER_RESP_SERVICE_CODE, r.Status)

	return r
}

// RespError 错误返回
func RespError(ctx *gin.Context, err xerr.XErr) *TrackedErrorResponse {
	logID := GetLogIDFromCtx(ctx)
	status := err.Code()

	r := &TrackedErrorResponse{
		Response: Response{
			Status: status,
			ErrMsg: GetMsg(status),
			Data:   "",
		},
		LogID: logID,
	}

	ctx.Set(consts.CONTEXT_HEADER_RESP_BODY, r.Data)
	ctx.Set(consts.CONTEXT_HEADER_RESP_MSG, r.ErrMsg)
	ctx.Set(consts.CONTEXT_HEADER_RESP_HTTP_CODE, http.StatusOK)
	ctx.Set(consts.CONTEXT_HEADER_RESP_SERVICE_CODE, r.Status)
	// todo add error

	return r
}

func GetLogIDFromCtx(ctx *gin.Context) (logID string) {
	spanCtxInterface, _ := ctx.Get(consts.CONTEXT_HEADER_CONTEXT_SPAN)
	str := fmt.Sprintf("%v", spanCtxInterface)
	re := regexp.MustCompile(`([0-9a-fA-F]{16})`)

	match := re.FindStringSubmatch(str)
	if len(match) > 0 {
		return match[1]
	}
	return ""
}
