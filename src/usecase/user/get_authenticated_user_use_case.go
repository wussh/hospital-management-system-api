package user

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
)

type getAuthenticatedUserUseCase struct {
	userRepository  domain.UserRepository
	jwtTokenManager application.TokenManager
}

func NewGetAuthenticatedUserUseCase(
	userRepository domain.UserRepository,
	jwtTokenManager application.TokenManager,
) domain.GetAuthenticatedUserUseCase {
	return &getAuthenticatedUserUseCase{
		userRepository:  userRepository,
		jwtTokenManager: jwtTokenManager,
	}
}

func (u *getAuthenticatedUserUseCase) Execute(
	authorizationHeader entity.AuthorizationHeader,
) (entity.User, int, error) {
	decodedPayload, code, err := u.jwtTokenManager.DecodeAccessTokenPayload(authorizationHeader.AccessToken)
	if err != nil {
		return entity.User{}, code, err
	}

	return u.userRepository.GetUserByID(decodedPayload.ID)
}
