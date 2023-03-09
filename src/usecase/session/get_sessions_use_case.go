package session

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
)

type getSessionsUseCase struct {
	sessionRepository domain.SessionRepository
	userRepository    domain.UserRepository
	jwtTokenManager   application.TokenManager
}

func NewGetSessionsUseCase(
	sessionRepository domain.SessionRepository,
	userRepository domain.UserRepository,
	jwtTokenManager application.TokenManager,
) domain.GetSessionsUseCase {
	return &getSessionsUseCase{
		sessionRepository: sessionRepository,
		userRepository:    userRepository,
		jwtTokenManager:   jwtTokenManager,
	}
}

func (u *getSessionsUseCase) Execute(
	payload entity.GetSessionParams,
	authorizationHeader entity.AuthorizationHeader,
) ([]entity.Session, int, error) {
	decodedPayload, code, err := u.jwtTokenManager.DecodeAccessTokenPayload(authorizationHeader.AccessToken)
	if err != nil {
		return []entity.Session{}, code, err
	}

	user, code, err := u.userRepository.GetUserByID(decodedPayload.ID)
	if err != nil {
		return []entity.Session{}, code, err
	}

	if user.Role == "Admin" {
		return u.sessionRepository.GetSessions()
	}

	switch payload.Status {
	case "completed":
		return u.sessionRepository.GetCompletedSessionsByDoctorID(decodedPayload.ID)
	case "cancelled":
		return u.sessionRepository.GetCancelledSessionsByDoctorID(decodedPayload.ID)
	case "queued":
		return u.sessionRepository.GetQueuedSessionsByDoctorID(decodedPayload.ID)
	}

	return u.sessionRepository.GetSessionsByDoctorID(decodedPayload.ID)
}
