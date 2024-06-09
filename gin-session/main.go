package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 创建一个 cookie session 中间件
	store := cookie.NewStore([]byte("secret"))

	// 使用 sessions.Sessions() 函数将 session 中间件绑定到 gin 上
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/hello", func(c *gin.Context) {
		// 从 session 中获取 hello 值
		session := sessions.Default(c)

		// 如果 hello 值不存在，则设置默认值
		if session.Get("hello") != "world" {
			session.Set("hello", "world")
			err := session.Save()
			if err != nil {
				return
			}
		}

		c.JSON(200, gin.H{"hello": session.Get("hello")})
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}

	fmt.Println("Server started on port 8080")

}
