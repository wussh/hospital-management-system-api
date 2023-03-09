package application

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type TokenManager interface {
	GenerateRefreshToken(payload entity.AuthenticationPayload) (string, int, error)
	GenerateAccessToken(payload entity.AuthenticationPayload) (string, int, error)
	VerifyRefreshToken(refreshToken string) (int, error)
	DecodeRefreshTokenPayload(refreshToken string) (entity.AuthenticationPayload, int, error)
	DecodeAccessTokenPayload(accessToken string) (entity.AuthenticationPayload, int, error)
}
