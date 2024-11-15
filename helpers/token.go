package helpers

import (
	"golang-backend/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var mySigningKey = []byte("secret")

type MyCustomClaims struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CreateToken(user *models.User) (string, error) {
	claims := MyCustomClaims {
		int(user.ID),
		user.Name,
		user.Email,
		jwt.RegisteredClaims {
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)), // 1 hour
			IssuedAt: jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString(mySigningKey)

	return ss, err
}