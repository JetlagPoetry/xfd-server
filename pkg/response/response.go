package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"xfd-backend/pkg/consts"
)

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
	LogID  string      `json:"log_id"`
}

// TrackedErrorResponse 有追踪信息的错误反应
type TrackedErrorResponse struct {
	Response
	LogID string `json:"log_id"`
}

// RespSuccess 带data成功返回
func RespSuccess(ctx *gin.Context, data interface{}, code ...int) *Response {
	logID := GetLogIDFromCtx(ctx)
	status := SUCCESS
	if code != nil {
		status = code[0]
	}

	if data == nil {
		data = "操作成功"
	}

	r := &Response{
		Status: status,
		Data:   data,
		Msg:    GetMsg(status),
		LogID:  logID,
	}
	ctx.Set(consts.CONTEXT_HEADER_RESP_BODY, r.Data)
	ctx.Set(consts.CONTEXT_HEADER_RESP_MSG, r.Msg)
	ctx.Set(consts.CONTEXT_HEADER_RESP_HTTP_CODE, http.StatusOK)
	ctx.Set(consts.CONTEXT_HEADER_RESP_SERVICE_CODE, r.Status)

	return r
}

// RespError 错误返回
// todo 修改错误返回
func RespError(ctx *gin.Context, err error, data string, code ...int) *TrackedErrorResponse {
	logID := GetLogIDFromCtx(ctx)
	status := ERROR
	if code != nil {
		status = code[0]
	}

	r := &TrackedErrorResponse{
		Response: Response{
			Status: status,
			Msg:    GetMsg(status),
			Data:   data,
			Error:  err.Error(),
		},
		LogID: logID,
	}

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
