package authentication

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type userLogoutUseCase struct {
	authenticationRepository domain.AuthenticationRepository
}

func NewUserLogoutUseCase(authenticationRepository domain.AuthenticationRepository) domain.StaffLogoutUseCase {
	return &userLogoutUseCase{
		authenticationRepository: authenticationRepository,
	}
}

func (u *userLogoutUseCase) Execute(payload entity.RefreshTokenPayload) (int, error) {
	if code, err := u.authenticationRepository.VerifyRefreshTokenExistence(payload); err != nil {
		return code, err
	}

	return u.authenticationRepository.DeleteRefreshToken(payload)
}
