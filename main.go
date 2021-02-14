package main

import (
	"github.com/gin-gonic/gin"
)

// ためし: https://gin-gonic.com/ja/docs/quickstart/
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.Run() // 8080 でサーバー立てる
}
