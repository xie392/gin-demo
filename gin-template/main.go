package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 设置模板目录
	r.LoadHTMLGlob("templates/*")

	// 设置静态文件目录
	r.Static("/assets", "./assets")

	// 定义路由和处理函数
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "gin-template-demo",
		})
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}

	fmt.Println("Server is running on port 8080")
}
