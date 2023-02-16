package jwt

import (
	"douyin-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// JWT define jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = 0 // success
		token := c.Query("token")
		if token == "" {
			code = -1 // fail
		} else {
			// parse token
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = -1 // fail to parse token
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = -1 // token expired
			}
		}
		if code != 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  "token not exists or expired",
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
