package sazs

import (
	"github.com/zeropen/pkg/types"
)

type Config struct {
	UserQuery UserQuery
	OTPQuery  OTPQuery
}

type DB interface {
	Connect() (Config, error)
}

type UserQuery interface {
	GetUser(eadd string) (*types.User, error)
	// InsertOTP(eadd string, otp string, now int64) error
	// GetOTPs(eadd string, since int64) ([]types.OTP, error)
}

type OTPQuery interface {
	GetOTPs(email string, since int64) (*[]types.OTP, error)
	InsertOTP(email string, otp string) error
}
