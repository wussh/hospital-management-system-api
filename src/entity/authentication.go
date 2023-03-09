package entity

type Authentication struct {
	Token string `gorm:"not null"`
}

type RefreshTokenPayload struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type NewLogin struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	UserName     string `json:"userName"`
}

type NewAccessToken struct {
	AccessToken string `json:"accessToken"`
}

type AuthorizationHeader struct {
	AccessToken string
}
