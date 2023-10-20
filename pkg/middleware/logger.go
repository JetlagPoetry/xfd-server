package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"strings"
	"time"
	"xfd-backend/pkg/consts"
	"xfd-backend/pkg/response"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		req, _ := c.GetRawData()
		reqID := response.GetLogIDFromCtx(c)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(req))

		reqTemp := req
		file, _, _ := c.Request.FormFile("file")
		if file != nil {
			reqTemp = []byte("")
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(req))
		log.Printf("[%s] %s | %s | requestBody=%s\n", reqID, c.Request.Method, c.Request.RequestURI, strings.TrimSpace(string(reqTemp)))

		c.Next()

		latency := time.Now().Sub(start)
		respBody, _ := c.Get(consts.CONTEXT_HEADER_RESP_BODY)
		log.Printf("[%s] %s | %s | responseBody=%s, latency=%s\n", reqID, c.Request.Method, c.Request.RequestURI, respBody, latency)
	}
}
