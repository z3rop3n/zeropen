package user

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/zeropen/app/spector/token"
	"github.com/zeropen/pkg/types"
	"github.com/zeropen/pkg/utils"
)

const (
	HOUR_1_IN_MILI    = token.HOUR_1_IN_MILI
	DAYS_30_IN_MILI   = token.DAYS_30_IN_MILI_SECONDS
	MINUTES_5_IN_MILI = token.MINUTES_5_IN_MILI
)

func (uApi *UserAPI) Signup(eadd string) (int, error) {

	otp := utils.GenerateOTP(6)
	err := utils.SendEmail(eadd, otp, uApi.appConfig.EmailFrom, uApi.appConfig.EmailAuthGoogle)
	if err != nil {
		return 500, err
	}
	uApi.config.OTPQuery.InsertOTP(eadd, otp)
	uApi.config.UserQuery.CreateOne(types.User{
		Email: &eadd,
	})

	return 200, nil
}

func (uApi *UserAPI) VerifyOTP(c context.Context, eadd string, otp string, deviceId string, platform string, location string) (int, *VerifyOTPResponse, error) {
	now := time.Now().UTC().UnixMilli()
	since := now - HOUR_1_IN_MILI
	otps, err := uApi.config.OTPQuery.GetOTPs(eadd, since)
	user, uerr := uApi.config.UserQuery.GetByEmailId(eadd)
	if uerr != nil {
		return 500, nil, uerr
	}

	if err != nil {
		return 500, nil, err
	}
	if len(*otps) == 0 {
		return 406, nil, fmt.Errorf("no OTPs found for email %s", eadd)
	}

	otpVerified := false
	for _, o := range *otps {
		if *o.OTP == otp {
			otpVerified = true
			break
		}
	}

	if !otpVerified {
		return 406, nil, fmt.Errorf("OTP not found")
	}
	tokenController := *uApi.tokenController

	refreshTokenExp := now + DAYS_30_IN_MILI
	refreshTokenId := uuid.New().String()
	refreshToken, err2 := tokenController.CreateRefreshToken(refreshTokenId, now, refreshTokenExp, *user.Id)
	accessToken, err3 := tokenController.CreateAccessToken(*user.Id, refreshTokenId, now+MINUTES_5_IN_MILI, now, now+HOUR_1_IN_MILI)

	if err3 != nil || err2 != nil {
		return 500, nil, fmt.Errorf("failed to create tokens")
	}

	uApi.config.RefreshTokenQuery.CreateOne(refreshTokenId, *refreshToken, refreshTokenExp, deviceId, platform, location, *user.Id)
	verifyOTPResponse := VerifyOTPResponse{
		AccessToken:  *accessToken,
		RefreshToken: *refreshToken,
	}
	return 200, &verifyOTPResponse, nil
}

// Get User Profile
func (uApi *UserAPI) GetUserProfile(c context.Context) (int, *types.User, error) {
	accessToken, ok := c.Value(token.AccessTokenKey).(types.AccessToken)
	if !ok {
		return 401, nil, fmt.Errorf("unauthorized")
	}
	user, err := uApi.config.UserQuery.GetById(accessToken.UserId)
	if err != nil {
		return 500, nil, err
	}
	return 200, user, nil
}
