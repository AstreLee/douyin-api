package jwt

import (
	"douyin-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// JWT 自定义中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = 0 // 获取成功
		token := c.Query("token")
		if token == "" {
			code = -1 // -1表示获取token失败
		} else {
			// 解析token
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = -1 // token解析失败
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = -1 // token已经失效
			}
		}
		if code != 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  "token不存在或失效",
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
