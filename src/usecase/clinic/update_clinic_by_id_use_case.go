package clinic

import (
	"fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
)

type updateClinicByIDUseCase struct {
	clinicRepository domain.ClinicRepository
	jwtTokenManager  application.TokenManager
	userRepository   domain.UserRepository
}

func NewUpdateClinicByIDUseCase(
	clinicRepository domain.ClinicRepository,
	jwtTokenManager application.TokenManager,
	userRepository domain.UserRepository,
) domain.UpdateClinicByIDUseCase {
	return &updateClinicByIDUseCase{
		clinicRepository: clinicRepository,
		jwtTokenManager:  jwtTokenManager,
		userRepository:   userRepository,
	}
}

func (u *updateClinicByIDUseCase) Execute(
	payload entity.UpdateClinicPayload,
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

	return u.clinicRepository.UpdateClinicByID(payload)
}
