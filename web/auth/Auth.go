package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type JwtCustomClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(email string) (string, error) {
	customClaim := &JwtCustomClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaim)

	// 署名付きトークンを生成
	signing_token, err := token.SignedString([]byte("secret"))
	if err != nil {
		return signing_token, err
	}

	return signing_token, nil
}

func CurrentUserEmail(c echo.Context) interface{} {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaim)
	return claims.Email
}
