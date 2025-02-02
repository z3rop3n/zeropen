package token

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/zeropen/pkg/types"

	t "github.com/zeropen/pkg/token"
)

const (
	DAYS_30_IN_MILI_SECONDS = 30 * 24 * 60 * 60 * 1000
	HOUR_1_IN_MILI          = 60 * 60 * 1000
	MINUTES_5_IN_MILI       = 60 * 1000
)

type TokenController interface {
	CreateRefreshToken(id string, iat int64, exp int64, userId string) (*string, error)
	CreateAccessToken(userId string, refreshTokenId string, whiteListedExp int64, iat int64, exp int64) (*string, error)
	AccessAuthMiddleware(h http.Handler) func(w http.ResponseWriter, r *http.Request)
}

type Token struct {
	JWT_SECRET_KEY string
}

type contextKey string

const AccessTokenKey contextKey = "access_token_obj"

func NewTokenObj(jwtSecretKey string) *Token {
	return &Token{
		JWT_SECRET_KEY: jwtSecretKey,
	}
}

func Create(userId string, deviceId string, platform string, location string, metadata map[string]string) {
	now := time.Now().UTC().UnixMilli()
	refreshToken := types.RefreshToken{
		UserId: userId,
		Exp:    now + DAYS_30_IN_MILI_SECONDS,
		Iat:    now,
	}

	var myMap map[string]interface{}
	data, _ := json.Marshal(refreshToken)
	json.Unmarshal(data, &myMap)
	t.Create(myMap, "")

}

func (tok Token) CreateRefreshToken(id string, iat int64, exp int64, userId string) (*string, error) {
	refreshToken := types.RefreshToken{
		UserId: userId,
		Iat:    iat,
		Exp:    exp,
		Id:     id,
	}
	var myMap map[string]interface{}
	data, _ := json.Marshal(refreshToken)
	json.Unmarshal(data, &myMap)
	return t.Create(myMap, tok.JWT_SECRET_KEY)
}

func (tok Token) CreateAccessToken(userId string, refreshTokenId string, whiteListedExp int64, iat int64, exp int64) (*string, error) {
	accessToken := types.AccessToken{
		UserId:         userId,
		RefreshTokenId: refreshTokenId,
		WhiteListedExp: whiteListedExp,
		Iat:            iat,
		Exp:            exp,
	}
	var myMap map[string]interface{}
	data, _ := json.Marshal(accessToken)
	json.Unmarshal(data, &myMap)
	return t.Create(myMap, tok.JWT_SECRET_KEY)
}

// Verify the access token
func (tok Token) VerifyAccessToken(accessToken string) (*types.AccessToken, error) {
	token, err := t.Verify[types.AccessToken](accessToken, tok.JWT_SECRET_KEY)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (tok Token) AccessAuthMiddleware(h http.Handler) func(w http.ResponseWriter, r *http.Request) {
	return (func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		tokenObj, err := tok.VerifyAccessToken(token)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), AccessTokenKey, *tokenObj)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (tApi *TokenAPI) RefreshAccessToken(refreshToken string) (int, *string, error) {
	token, err := t.Verify[types.RefreshToken](refreshToken, tApi.appConfig.JWT_AUTH_SECRET)
	now := time.Now().UTC().UnixMilli()
	if err != nil {
		return 401, nil, err
	}
	tokenObj, err := tApi.config.RefreshTokenQuery.FindOneById(token.Id)
	if err != nil {
		return 401, nil, err
	}
	if tokenObj.UserId != token.UserId {
		return 401, nil, fmt.Errorf("unauthorized")
	}
	if tokenObj.Exp < now {
		return 401, nil, fmt.Errorf("token expired")
	}
	if !tokenObj.IsActive {
		return 401, nil, fmt.Errorf("token is revoked")
	}

	newTok := NewTokenObj(tApi.appConfig.JWT_AUTH_SECRET)
	newAccessToken, err := newTok.CreateAccessToken(tokenObj.UserId, tokenObj.Id, now+MINUTES_5_IN_MILI, now, now+HOUR_1_IN_MILI)
	if err != nil {
		return 401, nil, err
	}

	return 200, newAccessToken, nil
}
