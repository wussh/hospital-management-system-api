package schedule

import (
	"fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
)

type deleteScheduleByIDUseCase struct {
	scheduleRepository domain.ScheduleRepository
	userRepository     domain.UserRepository
	jwtTokenManager    application.TokenManager
}

func NewDeleteScheduleByIDUseCase(
	scheduleRepository domain.ScheduleRepository,
	userRepository domain.UserRepository,
	jwtTokenManager application.TokenManager,
) domain.DeleteScheduleByIDUseCase {
	return &deleteScheduleByIDUseCase{
		scheduleRepository: scheduleRepository,
		userRepository:     userRepository,
		jwtTokenManager:    jwtTokenManager,
	}
}

func (u *deleteScheduleByIDUseCase) Execute(
	payload entity.ScheduleIDPayload,
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

	if decodedPayload.ID != payload.UserID {
		if user.Role != "Admin" && user.Role != "Staff" {
			return http.StatusForbidden, fmt.Errorf("restricted resource")
		}
	}
	schedule, code, err := u.scheduleRepository.GetScheduleByID(payload.ID)

	if *schedule.UserID != payload.UserID {
		if user.Role != "Admin" && user.Role != "Staff" {
			return http.StatusForbidden, fmt.Errorf("restricted resource")
		}
	}

	return u.scheduleRepository.DeleteScheduleByID(payload.ID)
}
