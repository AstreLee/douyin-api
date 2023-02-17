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

// FindByUsername find user according to username of user
func FindByUsername(username string) (common.User, int64) {
	var user common.User
	result := db.DB.Where("username = ?", username).Find(&user)
	return user, result.RowsAffected
}

// FindByID find user according to ID of user
func FindByID(ID uint) (common.User, int64) {
	user := common.User{}
	result := db.DB.Where("id = ?", ID).Find(&user)
	return user, result.RowsAffected
}

// SaveUser just save
func SaveUser(username, password string) (common.User, int64) {
	user := common.User{
		Username:        username,
		Password:        password,
		FollowCount:     0,
		FollowerCount:   0,
		Avatar:          "/static/default_avatar.jpg",
		BackgroundImage: "/static/default_background.png",
		Signature:       "这个用户很懒，什么简介都没有",
		TotalFavorited:  0,
		WorkCount:       0,
		FavoriteCount:   0,
	}
	result := db.DB.Omit("ID", "IsFollow").Create(&user)
	return user, result.RowsAffected
}
