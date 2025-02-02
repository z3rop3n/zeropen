package sazs

import (
	"github.com/zeropen/pkg/types"
)

type Config struct {
	UserQuery         UserQuery
	OTPQuery          OTPQuery
	RefreshTokenQuery RefreshTokenQuery
}

type DB interface {
	Connect() (Config, error)
}

type UserQuery interface {
	GetByEmailId(eadd string) (*types.User, error)
	GetById(id string) (*types.User, error)
	CreateOne(user types.User) error
}

type OTPQuery interface {
	GetOTPs(email string, since int64) (*[]types.OTP, error)
	InsertOTP(email string, otp string) error
}

type RefreshTokenQuery interface {
	CreateOne(refreshToken string, exp int64, deviceId string, platform string, location string, userId string) error
	FindAllByUserId(userId string) (*[]types.RefreshToken, error)
	FindOneById(id string) (*types.RefreshToken, error)
}
