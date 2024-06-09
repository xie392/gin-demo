package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 接收参数
	// http://127.0.0.1:8080/hello/world
	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "Hello " + name,
		})
	})

	// 接收多个参数
	// http://127.0.0.1:8080/sum/1/2?sum=3
	r.GET("/sum/:a/:b", func(c *gin.Context) {
		a := c.Param("a")
		b := c.Param("b")
		sum := c.Query("sum")
		c.JSON(200, gin.H{
			"message": "The sum of " + a + " and " + b + " is " + sum,
		})
	})

	// 接收 JSON 参数
	// curl -X POST http://127.0.0.1:8080/json -d '{"name": "Alice", "age": 25}' -H "Content-Type: application/json"
	r.POST("/json", func(c *gin.Context) {
		var json struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		// BindJSON 绑定 JSON 数据到结构体
		if c.ShouldBindJSON(&json) == nil {
			c.JSON(200, gin.H{
				"message": "Hello " + json.Name + ", you are " + fmt.Sprint(json.Age),
			})
		}
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
	fmt.Println("Server is running on port 8080")
}
