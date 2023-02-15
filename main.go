package main

import (
	"douyin-api/db"
	"douyin-api/service"
)

func main() {
	go service.RunMessageServer()
	initRouter(r)
	// 初始化数据库连接s
	err := db.InitGormDB()
	if err != nil {
		return
	}
	r.Run(":8080")
}
