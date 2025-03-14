package token

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zeropen/pkg/types"
)

func Create(mp map[string]interface{}, secretKey string) (*string, error) {
	now := time.Now()
	mp["iat"] = now.UTC().UnixMilli()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims(mp))

	byteSecretKey := []byte(secretKey)
	tokenString, err := token.SignedString(byteSecretKey)
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

// func Verify(token string, secretKey string) (*types.AccessToken, error) {
// 	var claims types.AccessToken
// 	tok, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(secretKey), nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	if !tok.Valid {
// 		return nil, fmt.Errorf("invalid token")
// 	}
// 	claimsmp := tok.Claims.(jwt.MapClaims)
// 	data, err := json.Marshal(claimsmp)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if err := json.Unmarshal(data, &claims); err != nil {
// 		return nil, err
// 	}
// 	fmt.Println("claims", claims)
// 	return &claims, nil
// }

type TokenType interface {
	types.RefreshToken | types.AccessToken
	GetExp() int64
}

func Verify[T TokenType](token string, secretKey string) (*T, error) {
	now := time.Now().UnixMilli()
	var claims T
	tok, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !tok.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claimsmp := tok.Claims.(jwt.MapClaims)
	data, err := json.Marshal(claimsmp)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &claims); err != nil {
		return nil, err
	}
	if now > claims.GetExp() {
		return nil, fmt.Errorf("token expired")
	}
	return &claims, nil
}
