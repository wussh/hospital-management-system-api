package user

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
)

type addUserUseCase struct {
	userRepository     domain.UserRepository
	bcryptPasswordHash application.PasswordHash
	jwtTokenManager    application.TokenManager
	nanoidIDGenerator  application.IDGenerator
}

func NewAddUserUseCase(
	userRepository domain.UserRepository,
	bcryptPasswordHash application.PasswordHash,
	jwtTokenManager application.TokenManager,
	nanoidIDGenerator application.IDGenerator,
) domain.AddUserUseCase {
	return &addUserUseCase{
		userRepository:     userRepository,
		bcryptPasswordHash: bcryptPasswordHash,
		jwtTokenManager:    jwtTokenManager,
		nanoidIDGenerator:  nanoidIDGenerator,
	}
}

func (u *addUserUseCase) Execute(
	payload entity.User,
	authorizationHeader entity.AuthorizationHeader,
) (entity.AddedUser, int, error) {
	decodedPayload, code, err := u.jwtTokenManager.DecodeAccessTokenPayload(authorizationHeader.AccessToken)
	if err != nil {
		return entity.AddedUser{}, code, err
	}

	user, code, err := u.userRepository.GetUserByID(decodedPayload.ID)
	if err != nil {
		return entity.AddedUser{}, code, err
	}

	if user.Role != "Admin" {
		return entity.AddedUser{}, http.StatusForbidden, fmt.Errorf("restricted resource")
	}

	if payload.Role == "Admin" {
		return entity.AddedUser{}, http.StatusBadRequest, fmt.Errorf("admin user already exists")
	}

	if payload.Role != "Doctor" && payload.Role == "Staff" {
		return entity.AddedUser{}, http.StatusBadRequest, fmt.Errorf("role must be Doctor or Staff")
	}

	if payload.Role == "Doctor" {
		if len(payload.License) == 0 {
			return entity.AddedUser{}, http.StatusBadRequest, fmt.Errorf("license can't be empty for Doctor role")
		}
		if len(payload.Speciality) == 0 {
			return entity.AddedUser{}, http.StatusBadRequest, fmt.Errorf("speciality can't be empty for Doctor role")
		}
	}

	generatedID, code, err := u.nanoidIDGenerator.Generate()
	if err != nil {
		return entity.AddedUser{}, code, err
	}
	payload.ID = fmt.Sprintf("%s-%s", strings.ToLower(payload.Role), generatedID)

	hashedPassword, code, err := u.bcryptPasswordHash.Hash(payload.Password)
	if err != nil {
		return entity.AddedUser{}, code, err
	}
	payload.Password = hashedPassword

	return u.userRepository.AddUser(payload)
}
