package postgresdriver

import (
	"time"

	"github.com/google/uuid"
	"github.com/zeropen/pkg/types"
	"gorm.io/gorm"
)

type PostgresRefreshTokenQuery struct {
	db *gorm.DB
}

type RefreshTokenModel struct {
	Id       string `json:"id"`
	Token    string `json:"token"`
	Exp      int64  `json:"exp"`
	Iat      int64  `json:"iat"`
	DeviceId string `json:"deviceId"`
	Platform string `json:"platform"`
	Location string `json:"location"`
	UserId   string `json:"userId"`
	IsActive bool   `json:"isActive"`
}

func NewPostgresRefreshTokenQuery(db *gorm.DB) *PostgresRefreshTokenQuery {
	db.AutoMigrate(&RefreshTokenModel{})
	return &PostgresRefreshTokenQuery{db: db}
}

func (p *PostgresRefreshTokenQuery) CreateOne(refreshToken string, exp int64, deviceId string, platform string, location string, userId string) error {
	id := uuid.New().String()
	iat := time.Now().UnixMilli()
	token := RefreshTokenModel{
		Id:     id,
		Token:  refreshToken,
		Exp:    exp,
		UserId: userId,
		Iat:    iat,
	}
	return p.db.Create(&token).Error
}

func (p *PostgresRefreshTokenQuery) FindAllByUserId(userId string) (*[]types.RefreshToken, error) {
	var tokens []RefreshTokenModel
	var refreshTokens []types.RefreshToken
	err := p.db.Where("user_id = ?", userId).Find(&tokens).Error
	for _, token := range tokens {
		refreshTokens = append(refreshTokens, types.RefreshToken{
			UserId: token.UserId,
			Iat:    token.Iat,
			Id:     token.Id,
			Exp:    token.Exp,
		})
	}
	return &refreshTokens, err
}

func (p *PostgresRefreshTokenQuery) FindOneById(id string) (*types.RefreshToken, error) {
	var token RefreshTokenModel
	err := p.db.Where("id = ?", id).Find(&token).Error
	return &types.RefreshToken{
		Id:     token.Id,
		UserId: token.UserId,
		Iat:    token.Iat,
		Exp:    token.Exp,
	}, err
}
