package dao

import (
	"douyin-api/db"
	"douyin-api/entity/common"
)

// IsFollow check if username1 follows username2
func IsFollow(userId1, userId2 uint) (isExist bool) {
	result := db.DB.Model(&common.RelationUserUser{
		FavoriteId:  userId1,
		FavoritedId: userId2,
	}).Find(&common.RelationUserUser{})
	if result.RowsAffected > 0 {
		isExist = true
		return
	} else {
		isExist = false
		return
	}
}
