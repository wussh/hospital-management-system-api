package user

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type getUserByIDUseCase struct {
	userRepository domain.UserRepository
}

func NewGetUserByIDUseCase(
	userRepository domain.UserRepository,
) domain.GetUserByIDUseCase {
	return &getUserByIDUseCase{
		userRepository: userRepository,
	}
}

func (u *getUserByIDUseCase) Execute(payload entity.UserIDPayload) (entity.User, int, error) {
	return u.userRepository.GetUserByID(payload.ID)
}
