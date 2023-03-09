package schedule

import (
	"fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
)

type getSchedulesByDoctorIDUseCase struct {
	scheduleRepository domain.ScheduleRepository
	userRepository     domain.UserRepository
	jwtTokenManager    application.TokenManager
}

func NewGetSchedulesByDoctorIDUseCase(
	scheduleRepository domain.ScheduleRepository,
	userRepository domain.UserRepository,
	jwtTokenManager application.TokenManager,
) domain.GetSchedulesByDoctorIDUseCase {
	return &getSchedulesByDoctorIDUseCase{
		scheduleRepository: scheduleRepository,
		userRepository:     userRepository,
		jwtTokenManager:    jwtTokenManager,
	}
}

func (u *getSchedulesByDoctorIDUseCase) Execute(
	payload entity.UserIDPayload,
	authorizationHeader entity.AuthorizationHeader,
) ([]entity.Schedule, int, error) {
	decodedPayload, code, err := u.jwtTokenManager.DecodeAccessTokenPayload(authorizationHeader.AccessToken)
	if err != nil {
		return []entity.Schedule{}, code, err
	}

	user, code, err := u.userRepository.GetUserByID(decodedPayload.ID)
	if err != nil {
		return []entity.Schedule{}, code, err
	}

	if decodedPayload.ID != payload.ID {
		if user.Role != "Admin" && user.Role != "Staff" {
			return []entity.Schedule{}, http.StatusForbidden, fmt.Errorf("restricted resource")
		}
	}

	return u.scheduleRepository.GetSchedulesByDoctorID(payload.ID)
}
