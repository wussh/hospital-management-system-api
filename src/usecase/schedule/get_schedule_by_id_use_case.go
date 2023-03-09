package schedule

import (
	"fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
)

type getScheduleByIDUseCase struct {
	scheduleRepository domain.ScheduleRepository
	userRepository     domain.UserRepository
	jwtTokenManager    application.TokenManager
}

func NewGetScheduleByIDUseCase(
	scheduleRepository domain.ScheduleRepository,
	userRepository domain.UserRepository,
	jwtTokenManager application.TokenManager,
) domain.GetScheduleByIDUseCase {
	return &getScheduleByIDUseCase{
		scheduleRepository: scheduleRepository,
		userRepository:     userRepository,
		jwtTokenManager:    jwtTokenManager,
	}
}

func (u *getScheduleByIDUseCase) Execute(
	payload entity.ScheduleIDPayload,
	authorizationHeader entity.AuthorizationHeader,
) (entity.Schedule, int, error) {
	decodedPayload, code, err := u.jwtTokenManager.DecodeAccessTokenPayload(authorizationHeader.AccessToken)
	if err != nil {
		return entity.Schedule{}, code, err
	}

	user, code, err := u.userRepository.GetUserByID(decodedPayload.ID)
	if err != nil {
		return entity.Schedule{}, code, err
	}

	if decodedPayload.ID != payload.UserID {
		if user.Role != "Admin" && user.Role != "Staff" {
			return entity.Schedule{}, http.StatusForbidden, fmt.Errorf("restricted resource")
		}
	}
	schedule, code, err := u.scheduleRepository.GetScheduleByID(payload.ID)

	if *schedule.UserID != payload.UserID {
		if user.Role != "Admin" && user.Role != "Staff" {
			return entity.Schedule{}, http.StatusForbidden, fmt.Errorf("restricted resource")
		}
	}

	return schedule, code, err
}
