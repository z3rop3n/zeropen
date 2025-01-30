package postgresdriver

import (
	"github.com/zeropen/pkg/types"
	"gorm.io/gorm"
)

type PostgresUserQuery struct {
	db *gorm.DB
}

type UserModel struct {
	Id            *string `json:"id"`
	Email         *string `json:"email"`
	FirstName     *string `json:"firstName"`
	LastName      *string `json:"lastName"`
	ProfilePicUrl *string `json:"profilePic"`
	Iat           *int64  `json:"iat"`
	IsVerified    *bool   `json:"isVerified"`
	IsDeleted     *bool   `json:"isDeleted"`
}

func NewPostgresUserQuery(db *gorm.DB) *PostgresUserQuery {
	db.AutoMigrate(&UserModel{})
	return &PostgresUserQuery{db: db}
}

func (p PostgresUserQuery) GetUser(eadd string) (*types.User, error) {
	var user UserModel
	if err := p.db.Where("email = ? AND is_deleted != ? AND is_verified != ?", eadd, true, true).First(&user).Error; err != nil {
		return nil, err
	}

	return &types.User{
		Id:            user.Id,
		Email:         user.Email,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		ProfilePicUrl: user.ProfilePicUrl,
	}, nil
}

// func (p PostgresUserQuery) GetOTPs(string, int64) ([]types.OTP, error) {
// 	var otps []types.OTP
// 	if err := p.db.Find(&otps).Error; err != nil {
// 		return otps, err
// 	} else {
// 		return []types.OTP{}, err
// 	}
// }

// func (p PostgresUserQuery) InsertOTP(string, string) error {
// 	return nil
// }
