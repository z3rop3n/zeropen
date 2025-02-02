package postgresdriver

import (
	"github.com/google/uuid"
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

func (p PostgresUserQuery) GetByEmailId(eadd string) (*types.User, error) {
	var user UserModel
	if err := p.db.Where("email = ? AND is_deleted IS NULL OR NOT ?", eadd, true).First(&user).Error; err != nil {
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

func (p PostgresUserQuery) GetById(id string) (*types.User, error) {
	var user UserModel
	if err := p.db.Where("id = ? AND is_deleted IS NULL OR NOT ?", id, true).First(&user).Error; err != nil {
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

func (p PostgresUserQuery) CreateOne(user types.User) error {
	id := uuid.New().String()
	userModel := UserModel{
		Id:            &id,
		Email:         user.Email,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		ProfilePicUrl: user.ProfilePicUrl,
	}
	return p.db.Create(&userModel).Error
}
