package user

import (
	"fmt"
	"time"

	"github.com/zeropen/pkg/utils"
)

const (
	HOUR_1_AGO_MILI = 60 * 60 * 1000
)

func (uApi *UserAPI) Signup(eadd string) (int, error) {
	_, uerr := uApi.config.UserQuery.GetUser(eadd)
	if uerr == nil {
		return 406, uerr
	}

	otp := utils.GenerateOTP(6)
	err := utils.SendEmail(eadd, otp, uApi.appConfig.EmailFrom, uApi.appConfig.EmailAuthGoogle)
	if err != nil {
		return 500, err
	}
	uApi.config.OTPQuery.InsertOTP(eadd, otp)

	return 200, nil
}

func (uApi *UserAPI) VerifyOTP(eadd string, otp string) (int, error) {
	now := time.Now().UTC().UnixMilli()
	since := now - HOUR_1_AGO_MILI
	otps, err := uApi.config.OTPQuery.GetOTPs(eadd, since)

	if err != nil {
		return 500, err
	}
	if len(*otps) == 0 {
		return 406, fmt.Errorf("no OTPs found for email %s", eadd)
	}

	for _, o := range *otps {
		if *o.OTP == otp {
			return 200, nil
		}
	}

	return 406, fmt.Errorf("invalid OTP")
}
