package dao

import (
	"douyin-api/db"
	"fmt"
	"testing"
)

func Test_findByUsernameAndPassword(t *testing.T) {
	err := db.InitGormDB()
	if err != nil {
		return
	}
	_, row := FindByUsernameAndPassword("Andys", "123456")
	fmt.Println(row)
}
