package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ためし: https://gin-gonic.com/ja/docs/quickstart/
// https://qiita.com/daitasu/items/9428220810816972b5d5
func main() {
	r := gin.Default()
	// $ curl localhost:8080
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	// $ curl localhost:8080/data
	r.GET("/data", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "Go言語",
			"tag":  "<go>",
		}
		// {"lang":"Go\u8a00\u8a9e","tag":"\u003cgo\u003e"} を返す
		c.AsciiJSON(http.StatusOK, data)
	})
	r.Run(":8080") // 8080 でサーバー立てる
}
