package dao

import (
	"douyin-api/db"
	"douyin-api/entity/common"
)

func FindByLatestTime(latestTime int64) ([]common.Video, int64) {
	var videoList []common.Video
	result := db.DB.Raw("SELECT * FROM video v INNER JOIN user u ON u.id = v.user_id WHERE latest_time > ? ORDER BY latest_time LIMIT 30", latestTime).Scan(&videoList)

	return videoList, result.RowsAffected
}
