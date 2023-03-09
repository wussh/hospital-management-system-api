package schedule

import (
	"fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
)

type addScheduleUseCase struct {
	scheduleRepository domain.ScheduleRepository
	userRepository     domain.UserRepository
	jwtTokenManager    application.TokenManager
	nanoidIDGenerator  application.IDGenerator
}

func NewAddScheduleUseCase(
	scheduleRepository domain.ScheduleRepository,
	userRepository domain.UserRepository,
	jwtTokenManager application.TokenManager,
	nanoidIDGenerator application.IDGenerator,
) domain.AddScheduleUseCase {
	return &addScheduleUseCase{
		scheduleRepository: scheduleRepository,
		userRepository:     userRepository,
		jwtTokenManager:    jwtTokenManager,
		nanoidIDGenerator:  nanoidIDGenerator,
	}
}

func (u *addScheduleUseCase) Execute(
	payload entity.Schedule,
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

	if decodedPayload.ID != *payload.UserID {
		if user.Role != "Admin" && user.Role != "Staff" {
			return http.StatusForbidden, fmt.Errorf("restricted resource")
		}
	}

	generatedID, code, err := u.nanoidIDGenerator.Generate()
	if err != nil {
		return code, err
	}
	payload.ID = fmt.Sprintf("medical-record-%s", generatedID)

	return u.scheduleRepository.AddSchedule(payload)
}
