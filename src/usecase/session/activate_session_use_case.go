package session

import (
	"fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
)

type activateSessionUseCase struct {
	sessionRepository domain.SessionRepository
	userRepository    domain.UserRepository
	jwtTokenManager   application.TokenManager
}

func NewActivateSessionUseCase(
	sessionRepository domain.SessionRepository,
	userRepository domain.UserRepository,
	jwtTokenManager application.TokenManager,
) domain.ActivateSessionUseCase {
	return &activateSessionUseCase{
		sessionRepository: sessionRepository,
		userRepository:    userRepository,
		jwtTokenManager:   jwtTokenManager,
	}
}

func (u *activateSessionUseCase) Execute(
	sessionIDPayload entity.SessionIDPayload,
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

	session, code, err := u.sessionRepository.GetSessionByID(sessionIDPayload.ID)
	if err != nil {
		return code, err
	}

	if *session.DoctorID != user.ID {
		if user.Role != "Admin" && user.Role != "Staff" {
			return http.StatusForbidden, fmt.Errorf("restricted resource")
		}
	}

	if session.Status != "Dalam antrian" {
		return http.StatusBadRequest, fmt.Errorf("can't activate session which status isn't Dalam antrian")
	}

	if code, err := u.sessionRepository.VerifyNoActiveSession(); err != nil {
		return code, err
	}

	return u.sessionRepository.UpdateSessionStatusToActive(sessionIDPayload.ID)
}
