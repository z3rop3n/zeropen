package postgresdriver

import (
	"time"

	"github.com/zeropen/pkg/types"
	"gorm.io/gorm"
)

type PostgresOTPQuery struct {
	db *gorm.DB
}

type OTPModel struct {
	Email *string `json:"email"`
	OTP   *string `json:"otp"`
	Iat   *int64  `json:"iat"`
}

func NewPostgresOTPQuery(db *gorm.DB) *PostgresOTPQuery {
	db.AutoMigrate(&OTPModel{})
	return &PostgresOTPQuery{db: db}
}

func (p *PostgresOTPQuery) GetOTPs(email string, since int64) (*[]types.OTP, error) {
	var otps []OTPModel
	if err := p.db.Where("email = ? AND Iat >= ?", email, since).Find(&otps).Error; err != nil {
		return nil, err
	}

	var otpsr []types.OTP
	for _, otp := range otps {
		otpsr = append(otpsr, types.OTP{
			Email: otp.Email,
			OTP:   otp.OTP,
		})
	}

	return &otpsr, nil
}

func (p *PostgresOTPQuery) InsertOTP(email string, otp string) error {
	now := time.Now().UTC().UnixMilli()
	otpModel := OTPModel{
		Email: &email,
		OTP:   &otp,
		Iat:   &now,
	}
	return p.db.Create(&otpModel).Error
}
