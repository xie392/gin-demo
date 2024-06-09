package main

import (
	"gin-mysql/database"
	"gin-mysql/routers"
)

func main() {
	// 初始化数据库
	database.InitDB()

	// 设置路由
	r := routers.SetupRouter()

	// 启动服务器
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
