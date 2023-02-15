package utils

import (
	"github.com/golang-jwt/jwt"
	"time"
)

// jwtSecret 签名密钥
var jwtSecret = "xin"

// Claims Claims保存了用户名和密码还有一些token中的标准配置
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken 根据用户名和密码生成token
// params：username、password：用户名和密码
// returns：token, error：token/nil 或 空/err
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Minute * 30).Unix() // 设置token30分钟后过期

	claims := Claims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,        // token过期时间
			IssuedAt:  time.Now().Unix(), // token发行时间
			Issuer:    "xin",             // token签发着
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken 根据Token字符串解析出Claims，进行获取用户名和密码
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
