package routerForward

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)

/*
	reqForward 请求转发函数
		isSSL: 是否使用ssl协议转发
		host: 转发目的host, example: 192.168.0.13:8080
		url: 转发目的url, example: /v1.01/test
*/
func reqForward(c *gin.Context, isSSL bool, host string, url string) {
	req := c.Request

	var reqScheme string
	if isSSL {
		// ssl协议转发
		reqScheme = "https"
	} else {
		reqScheme = "http"
	}

	// 转发路由的配置
	director := func(req *http.Request) {
		req.URL.Scheme = reqScheme
		req.URL.Host = host
		req.URL.Path = url
		req.Host = host
	}

	// 这里我们只转发请求, 返回不做额外处理. 目的服务的返回
	proxy := httputil.ReverseProxy{
		Director: director,
	}

	proxy.ServeHTTP(c.Writer, req)
}
