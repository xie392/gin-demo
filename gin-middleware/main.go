package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 在这里可以进行一些处理，比如打印日志
		fmt.Println("Before request")

		// 调用下一个处理器
		c.Next()

		// 在这里可以对响应进行一些处理
		fmt.Println("After request")
	}
}

func main() {
	r := gin.Default()

	// 全局加入中间件
	// 使用多个 r.Use(middleware(),middleware()...)
	r.Use(middleware())

	// 局部加 r.GET("/test",middleware(),func...)
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "test"})
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
	fmt.Println("Server is running on port 8080")
}
