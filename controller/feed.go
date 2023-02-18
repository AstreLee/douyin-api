package controller

import (
	"douyin-api/dao"
	"douyin-api/entity"
	"douyin-api/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// Feed get all video info
func Feed(c *gin.Context) {
	// get latest_time and token (if exist)
	latestTime := c.Query("latest_time")
	token := c.Query("token")

	// if latest_time != ""
	var timeGet int64
	if latestTime == "" {
		timeGet = time.Now().Unix()
	} else {
		timeGet, _ = strconv.ParseInt(latestTime, 10, 64)
	}
	videoList, row := dao.FindByLatestTime(timeGet)
	if row <= 0 {
		c.JSON(http.StatusOK, entity.VideoResponse{
			StatusCode: 0,
			StatusMsg:  "视频流获取失败",
			NextTime:   0,
			VideoList:  nil,
		})
	} else {
		var serverIP = "http://192.168.1.101:8080"
		for i := 0; i < len(videoList); i++ {
			if token != "" {
				claims, _ := utils.ParseToken(token)
				selfID := claims.ID
				if dao.IsFollow(selfID, videoList[i].Author.ID) {
					videoList[i].Author.IsFollow = true
				} else {
					videoList[i].Author.IsFollow = false
				}
				if dao.IsFavorite(selfID, videoList[i].ID) {
					videoList[i].IsFavorite = true
				} else {
					videoList[i].IsFavorite = false
				}
			}
			videoList[i].PlayUrl = serverIP + videoList[i].PlayUrl
			videoList[i].CoverUrl = serverIP + videoList[i].CoverUrl
		}

		fmt.Println(videoList)
		c.JSON(http.StatusOK, entity.VideoResponse{
			StatusCode: 0,
			StatusMsg:  "获取视频流成功",
			NextTime:   time.Now().Unix(),
			VideoList:  videoList,
		})
	}

}
