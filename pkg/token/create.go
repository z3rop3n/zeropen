package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Create(mp map[string]interface{}, secretKey string) (*string, error) {
	now := time.Now()
	mp["iat"] = now.UTC().UnixMilli()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims(mp))

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
