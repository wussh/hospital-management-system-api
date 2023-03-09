package user

import (
	"fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
)

type deleteUserAvatarUseCase struct {
	userRepository  domain.UserRepository
	jwtTokenManager application.TokenManager
}

func NewDeleteUserAvatarUseCase(
	userRepository domain.UserRepository,
	jwtTokenManager application.TokenManager,
) domain.DeleteUserAvatarUseCase {
	return &deleteUserAvatarUseCase{
		userRepository:  userRepository,
		jwtTokenManager: jwtTokenManager,
	}
}

func (u *deleteUserAvatarUseCase) Execute(
	payload entity.UserIDPayload,
	authorizationHeader entity.AuthorizationHeader,
) (int, error) {
	decodedPayload, code, err := u.jwtTokenManager.DecodeAccessTokenPayload(authorizationHeader.AccessToken)
	if err != nil {
		return code, err
	}

	user, code, err := u.userRepository.GetUserByID(decodedPayload.ID)
	if err != nil {
		return code, err
	}

	if user.ID != decodedPayload.ID {
		if user.Role != "Admin" {
			return http.StatusForbidden, fmt.Errorf("restricted resource")
		}
	}

	return u.userRepository.DeleteUserAvatar(payload)
}
