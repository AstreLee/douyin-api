package dao

import (
	"douyin-api/db"
	"douyin-api/entity/common"
)

func IsFavorite(userId, videoId uint) bool {
	var relationUserVideo common.RelationUserVideo
	result := db.DB.Where("user_id = ? AND video_id = ?", userId, videoId).Find(&relationUserVideo)
	if result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}
