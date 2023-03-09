package clinic

import (
	"fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
)

type deleteClinicByIDUseCase struct {
	clinicRepository domain.ClinicRepository
	jwtTokenManager  application.TokenManager
	userRepository   domain.UserRepository
}

func NewDeleteClinicByIDUseCase(
	clinicRepository domain.ClinicRepository,
	jwtTokenManager application.TokenManager,
	userRepository domain.UserRepository,
) domain.DeleteClinicByIDUseCase {
	return &deleteClinicByIDUseCase{
		clinicRepository: clinicRepository,
		jwtTokenManager:  jwtTokenManager,
		userRepository:   userRepository,
	}
}

func (u *deleteClinicByIDUseCase) Execute(
	id string,
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

	if user.Role != "Admin" {
		return http.StatusForbidden, fmt.Errorf("restricted resource")
	}

	return u.clinicRepository.DeleteClinicByID(id)
}
