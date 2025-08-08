package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/peopleig/food-ordering-go/pkg/config"
)

var jwt_secret = []byte(config.SECRET_KEY)

type Claims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(user_id int, role string) (string, error) {
	claims := &Claims{
		UserID: user_id,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwt_secret)
}
