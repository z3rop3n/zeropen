package token

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	t "github.com/zeropen/pkg/token"
)

const (
	DAYS_30_IN_SECONDS = 30 * 24 * 60 * 60
)

type RefreshToken struct {
	// jwt.RegisteredClaims
	UserId string `json:"userId"`
	Id     string `json:"id"`
	Iat    int64  `json:"iat"`
	Exp    int64  `json:"exp"`
	// DeviceId string `json:"deviceId"`
	// Platform string `json:"platform"`
	// Location string `json:"location"`
}

type AccessToken struct {
	// jwt.RegisteredClaims
	UserId         string `json:"userId"`
	RefreshTokenId string `json:"refreshToken"`
	WhiteListedExp int64  `json:"whiteListedExp"`
}

func Create(userId string, deviceId string, platform string, location string, metadata map[string]string) {
	refreshTokenId := uuid.New().String()
	now := time.Now().UTC().UnixMilli()
	refreshToken := RefreshToken{
		UserId: userId,
		Id:     refreshTokenId,
		Exp:    now + DAYS_30_IN_SECONDS,
		Iat:    now,
		// DeviceId: deviceId,
		// Platform: platform,
		// Location: location,
	}

	var myMap map[string]interface{}
	data, _ := json.Marshal(refreshToken)
	json.Unmarshal(data, &myMap)
	t.Create(myMap, "")

}
