package controller

import (
	"douyin-api/dao"
	"douyin-api/entity"
	"douyin-api/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Login(c *gin.Context) {
	// get username and password
	username := c.Query("username")
	password := c.Query("password")

	// call func to check whether user exists
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

func Register(c *gin.Context) {
	// get username and password
	username := c.Query("username")
	password := c.Query("password")

	// check if user is already exist
	_, row := dao.FindByUsernameAndPassword(username, password)
	// user exist, tip try again
	if row != 0 {
		c.JSON(http.StatusOK, entity.UserLoginResponse{
			StatusCode: -1,
			StatusMsg:  "user is already existed, please try again",
			UserId:     0,
			Token:      "",
		})
	} else {
		// user does not exist
		// call func to generate token according username and password
		// first save the user
		user, row := dao.SaveUser(username, password)
		if row == 0 {
			fmt.Println("fail to save user")
			return
		}
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

func UserInfo(c *gin.Context) {
	// get user_id and token
	userID := c.Query("user_id")

	// get info of user by ID
	id, _ := strconv.ParseUint(userID, 10, 64)
	user, row := dao.FindUserByID(uint(id))
	if row == 0 {
		fmt.Println("fail to get info of user！")
		return
	} else {
		c.JSON(http.StatusOK, user)
	}
}
