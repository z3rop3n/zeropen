package types

type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AccessToken struct {
	UserId         string `json:"userId"`
	RefreshTokenId string `json:"refreshToken"`
	WhiteListedExp int64  `json:"whiteListedExp"`
	Iat            int64  `json:"iat"`
	Exp            int64  `json:"exp"`
}

type RefreshToken struct {
	Id       string `json:"id"`
	UserId   string `json:"userId"`
	Iat      int64  `json:"iat"`
	Exp      int64  `json:"exp"`
	IsActive bool   `json:"isActive"`
}

func (rt RefreshToken) GetExp() int64 {
	return rt.Exp
}
func (at AccessToken) GetExp() int64 {
	return at.Exp
}
