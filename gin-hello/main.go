package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 单个路由
	// http://127.0.0.1:8080/ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 路由组
	v1 := r.Group("/api/v1")
	{
		// http://127.0.0.1:8080/api/v1/users
		v1.GET("/users", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "get users",
			})
		})
		// http://127.0.0.1:8080/api/v1/users
		v1.POST("/users", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "create user",
			})
		})
	}

	err := r.Run(":8080")
	if err != nil {
		return
	}
	fmt.Println("Server is running on port 8080")
}
