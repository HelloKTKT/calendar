package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var secretKey = []byte("gqv7jtu7VZD1dar")

func CreateToken(userID string) (string, error) {
	// 创建负载
	payload := jwt.MapClaims{
		"user":    userID,
		"exptime": time.Now().Add(time.Minute * 15).Unix(), // 令牌过期时间为15分钟
	}

	// 创建 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	// 签名并获取完整的 Token 字符串
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	// 解析 Token 字符串
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
