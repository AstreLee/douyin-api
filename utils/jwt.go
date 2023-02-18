package utils

import (
	"github.com/golang-jwt/jwt"
	"time"
)

// jwtSecret private key
var jwtSecret = "xin"

// Claims ID、username and some standard info
type Claims struct {
	ID       uint   `json:"Id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken generate token according username
func GenerateToken(username string, userId uint) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Minute * 30).Unix() // set the expired time to 30 minutes

	claims := Claims{
		ID:       userId,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,        // expired time
			IssuedAt:  time.Now().Unix(), // issue time
			Issuer:    "xin",             // issuer
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(jwtSecret))
	return token, err
}

// ParseToken parse token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
