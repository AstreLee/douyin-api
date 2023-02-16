package controller

import (
	"douyin-api/dao"
	"douyin-api/entity"
	"douyin-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	// get username and password
	username := c.Query("username")
	password := c.Query("password")

	// call func to check whether user is existed
	user, row := dao.FindByUsernameAndPassword(username, password)

	// row == 0 means the user does not exist
	if row == 0 {
		// set httpCode to 200 and the UserLoginResponse
		c.JSON(http.StatusOK, entity.UserLoginResponse{
			StatusCode: -1,
			StatusMsg:  "username or password wrong",
			UserId:     0,
			Token:      "",
		})
	} else {
		// call func to generate token according username and password
		token, err := utils.GenerateToken(username, password)
		if err == nil {
			// set httpCode to 200 and the UserLoginResponse
			c.JSON(http.StatusOK, entity.UserLoginResponse{
				StatusCode: 0,
				StatusMsg:  "login success",
				UserId:     user.ID,
				Token:      token,
			})
		}
	}
}
