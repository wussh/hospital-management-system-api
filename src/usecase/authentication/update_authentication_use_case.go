package authentication

import (
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
)

type updateAuthenticationUseCase struct {
	jwtTokenManager          application.TokenManager
	authenticationRepostiory domain.AuthenticationRepository
}

func NewUpdateAuthenticationUseCase(
	jwtTokenManager application.TokenManager,
	authenticationRepository domain.AuthenticationRepository,
) domain.UpdateAuthenticationUseCase {
	return &updateAuthenticationUseCase{
		jwtTokenManager:          jwtTokenManager,
		authenticationRepostiory: authenticationRepository,
	}
}

func (u *updateAuthenticationUseCase) Execute(payload entity.RefreshTokenPayload) (
	entity.NewAccessToken, int, error,
) {
	if code, err := u.jwtTokenManager.VerifyRefreshToken(payload.RefreshToken); err != nil {
		return entity.NewAccessToken{}, code, err
	}

	if code, err := u.authenticationRepostiory.VerifyRefreshTokenExistence(payload); err != nil {
		return entity.NewAccessToken{}, code, err
	}

	decodedPayload, code, err := u.jwtTokenManager.DecodeRefreshTokenPayload(payload.RefreshToken)
	if err != nil {
		return entity.NewAccessToken{}, code, err
	}

	accessToken, code, err := u.jwtTokenManager.GenerateAccessToken(decodedPayload)
	if err != nil {
		return entity.NewAccessToken{}, code, err
	}

	newAccessToken := entity.NewAccessToken{
		AccessToken: accessToken,
	}

	return newAccessToken, http.StatusOK, nil
}
