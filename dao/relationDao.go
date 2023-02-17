package dao

import (
	"douyin-api/db"
	"douyin-api/entity/common"
)

// IsFollow check if username1 follows username2
func IsFollow(username1, username2 string) (isExist bool) {
	result := db.DB.Model(&common.RelationShip{FavoriteName: username1, FavoritedName: username2}).Find(&common.RelationShip{})
	if result.RowsAffected > 0 {
		isExist = true
		return
	} else {
		isExist = false
		return
	}
}
