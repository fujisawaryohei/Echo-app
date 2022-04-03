package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

type IAuthenticator interface {
	GenerateToken(email string) (string, error)
}

type Authenticator struct {
	JwtCustomClaim
}

type JwtCustomClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func NewAuthenticator() *Authenticator {
	return &Authenticator{
		JwtCustomClaim: JwtCustomClaim{
			Email: "",
			StandardClaims: jwt.StandardClaims{
				IssuedAt:  time.Now().Unix(),
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		},
	}
}

func (a *Authenticator) GenerateToken(email string) (string, error) {
	a.JwtCustomClaim.Email = email
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a.JwtCustomClaim)

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

func CurrentUserEmail(c echo.Context) interface{} {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaim)
	return claims.Email
}
