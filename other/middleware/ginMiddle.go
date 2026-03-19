package main

import (
	"compress/gzip"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", c.GetHeader("Origin"))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	// 静态文件
	router.StaticFS("/", http.Dir(dir+"/"))
	// 使用中间件
	router.Use(CORSMiddleware())

	/*
		添加gzip支持(注意添加gzip 过滤)
			过滤路由: (regexp包不支持?! 零宽断言, 因此这里使用`^(/([^h]|[h][^t]|[h][t][^m]|[h][t][m][^l]|[h][t][m][l][^5])).*`
					过滤非 '/html5' 开头的路由. 等效于 ^(?!/html5).*")
					`^(/([^a]|[a][^p]|[a][p][^k])).*` 过滤非'/apk' 开头的路由
					gzip.WithExcludedPathsRegexs([]string{"^(?!/html5).*"}))
			过滤文件: 所有.pdf, .mp4, .png, .gif, .jpeg, .jpg的文件会过滤

	*/
	router.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPathsRegexs([]string{"^(/([^h]|[h][^t]|[h][t][^m]|[h][t][m][^l]|[h][t][m][l][^5])).*"}),
		gzip.WithExcludedExtensions([]string{".pdf", ".mp4", ".png", ".gif", ".jpeg", ".jpg"})))
}
