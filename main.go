package main

import "douyin-api/db"

func main() {
	// init database connection
	err := db.InitGormDB()
	if err != nil {
		return
	}
	initRouter()
}
