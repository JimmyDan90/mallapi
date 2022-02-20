package common

import (
	"github.com/golang-jwt/jwt"
	"mallapi/global"
	"time"
)

var SigningKey = []byte(global.Config.Jwt.SigningKey)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.Claims
}

// 生成Token
func GenerateToken(username, password string) (string, error) {
	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 60 * 60,
			Issuer: username,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(SigningKey)
}

// 验证Token
func VerifyToken(tokenString string) error {
	_, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return SigningKey, nil
	})
	return err
}


