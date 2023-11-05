package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"regexp"
	"strings"
	"xfd-backend/pkg/consts"
	"xfd-backend/pkg/xerr"
)

// Response 基础序列化器
type Response struct {
	LogID  string      `json:"log_id,omitempty"`
	Status xerr.XCode  `json:"status,omitempty"`
	Msg    string      `json:"msg,omitempty"`
	Data   interface{} `json:"data,omitempty"`
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
func RespError(ctx *gin.Context, err xerr.XErr) *Response {
	logID := GetLogIDFromCtx(ctx)
	if err == nil {
		err = xerr.DefaultXErr
	}
	status := err.Code()

	r := &Response{
		Status: status,
		Msg:    GetMsg(status),
		Data:   err.Error(),
		LogID:  logID,
	}

	ctx.Set(consts.CONTEXT_HEADER_RESP_BODY, r.Data)
	ctx.Set(consts.CONTEXT_HEADER_RESP_MSG, r.Msg)
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

func SetLogId() gin.HandlerFunc {
	return func(c *gin.Context) {
		logID := strings.Replace(uuid.New().String(), "-", "", -1)
		c.Set(consts.CONTEXT_HEADER_CONTEXT_SPAN, logID)
		c.Request.Header.Add(consts.XRequestIDHeader, logID)
		c.Writer.Header().Set(consts.XRequestIDHeader, logID)

		c.Next()
	}

}
