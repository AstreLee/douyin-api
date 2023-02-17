package entity

import "douyin-api/entity/common"

type UserInfo struct {
	StatusCode  uint   `json:"status_code"`
	StatusMsg   string `json:"status_msg,omitempty"`
	common.User `json:"user,omitempty"`
}
