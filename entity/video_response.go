package entity

import "douyin-api/entity/common"

type VideoResponse struct {
	StatusCode uint           `json:"status_code"`
	StatusMsg  string         `json:"status_msg,omitempty"`
	NextTime   int64          `json:"next_time,omitempty"`
	VideoList  []common.Video `json:"video_list,omitempty"`
}
