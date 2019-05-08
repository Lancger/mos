package util

import (
	"mos/src/pkg/setting"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string `json:username`
	jwt.StandardClaims
}

// GenerateToken 生成JWT Token
func GenerateToken(username string) (string, error) {
	nowTime := time.Now()
	expriedTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expriedTime.Unix(),
			Issuer:    "jwt-go",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken 解析Token
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
