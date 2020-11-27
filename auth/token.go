package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// TokenTimeout 访问服务token过期时间 3小时
const TokenTimeout time.Duration = 3 * time.Second

type Claims struct {
	jwt.StandardClaims
}

// generateToken 生成token
func generateToken() (string, error) {
	//设置token有效时间
	nowTime := time.Now()
	expireTime := nowTime.Add(TokenTimeout)
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "hjx-services",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte("hjx"))
	return token, err
}

func validateToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("hjx"), nil
	})
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return err
}
