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

// SaveUser just save
func SaveUser(username, password string) (common.User, int64) {
	user := common.User{Username: username, Password: password, FollowCount: 0, FollowerCount: 0}
	result := db.DB.Omit("ID").Create(&user)
	return user, result.RowsAffected
}

// FindUserByID find user according to ID of user
func FindUserByID(ID uint) (common.User, int64) {
	user := common.User{}
	result := db.DB.Where("id = ?", ID).Find(&user)
	return user, result.RowsAffected
}
