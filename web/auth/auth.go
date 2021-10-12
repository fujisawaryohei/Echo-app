package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

type JwtCustomClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func SignKey() ([]byte, error) {
	secretKey := make([]byte, 10)
	file, err := os.Open("secret.key")
	if err != nil {
		return []byte{}, err
	}
	defer file.Close()

	if _, err := file.Read(secretKey); err != nil {
		return []byte{}, nil
	}
	return secretKey, nil
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
	signKey, err := SignKey()
	if err != nil {
		return "", err
	}

	signing_token, err := token.SignedString(signKey)
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
