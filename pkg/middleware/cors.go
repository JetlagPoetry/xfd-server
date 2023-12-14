package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method // 请求方法
		//// origin := c.Request.Header.Get("Origin") //请求头部
		//var headerKeys []string // 声明请求头keys
		//for k := range c.Request.Header {
		//	headerKeys = append(headerKeys, k)
		//}
		//headerStr := strings.Join(headerKeys, ", ")
		//if headerStr != "" {
		//	headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		//} else {
		//	headerStr = "access-control-allow-origin, access-control-allow-headers"
		//}
		// 增加用户请求信息
		ip := c.Request.Header.Get("X-real-ip")
		if ip == "" {
			ip = c.ClientIP()
		}
		domain := c.Request.Host
		log.Printf("ip:[%v],domain:[%v]\n", ip, domain)
		//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// 允许访问所有域
		c.Header("Access-Control-Allow-Origin", "*")
		// 服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		// header的类型
		c.Header("Access-Control-Allow-Headers", "Sec-CH-UA, Sec-CH-UA-Mobile, Sec-CH-UA-Platform, X-XSRF-TOKEN, XSRF-TOKEN, Referer, Content-Type, Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-ResourceType, Pragma, accessToken, X-AppId, X-Signature, X-Signature-Method")
		// 允许跨域设置	 可以返回其他子段
		// 跨域关键设置 让浏览器可以解析
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-ResourceType,Expires,Last-Modified,Pragma,FooBar")
		// 缓存请求信息 单位为秒
		c.Header("Access-Control-Max-Age", "172800")
		//	跨域请求是否需要带cookie信息 默认设置为true
		c.Header("Access-Control-Allow-Credentials", "true")
		// 设置返回格式是json
		c.Set("content-type", "application/json")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //	处理请求
	}
}
