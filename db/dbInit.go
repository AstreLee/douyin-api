package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitGormDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/douyin_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("数据库连接失败：%v\n", err)
	} else {
		fmt.Printf("数据库连接成功\n")
		DB = db
	}
	return err
}
