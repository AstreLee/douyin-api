package dao

import (
	"douyin-api/db"
	"douyin-api/entity/common"
)

// FindByUsernameAndPassword find user according to username and password, return user(if exists) and the affectRows
func FindByUsernameAndPassword(username, password string) (common.User, int64) {
	var user common.User
	result := db.DB.Where("username = ? AND password = ?", username, password).Find(&user)
	return user, result.RowsAffected
}
