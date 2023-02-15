package controller

import (
	"github.com/gin-gonic/gin"
)

var userIdSequence = int64(1)

func Register(c *gin.Context) {
}

func Login(c *gin.Context) {
	// 获取用户名和密码
	//username := c.Query("username")
	//password := c.Query("password")

}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
}
