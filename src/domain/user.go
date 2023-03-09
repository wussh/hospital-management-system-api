package domain

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	PostUserHandler(c echo.Context) error
	PutUserAvatarHandler(c echo.Context) error
	DeleteUserAvatarHandler(c echo.Context) error
	GetUserByIDHandler(c echo.Context) error
	GetAuthenticatedUserHandler(c echo.Context) error
}

type UserRepository interface {
	AddUser(payload entity.User) (entity.AddedUser, int, error)
	GetUserByID(id string) (entity.User, int, error)
	UpdateUserAvatar(payload entity.UpdateAvatarLocationPayload) (int, error)
	DeleteUserAvatar(payload entity.UserIDPayload) (int, error)
	GetUserDoctorByID(id string) (entity.User, int, error)
	GetUserStaffByID(id string) (entity.User, int, error)
	GetUserDoctorsByClinicID(clinicID string) ([]entity.User, int, error)
}

type AddUserUseCase interface {
	Execute(payload entity.User, authorizationHeader entity.AuthorizationHeader) (entity.AddedUser, int, error)
}

type UpdateUserAvatarUseCase interface {
	Execute(
		payload entity.UpdateAvatarPayload,
		authorizationHeader entity.AuthorizationHeader,
	) (entity.UpdatedAvatar, int, error)
}

type DeleteUserAvatarUseCase interface {
	Execute(
		payload entity.UserIDPayload,
		authorizationHeader entity.AuthorizationHeader,
	) (int, error)
}

type GetUserByIDUseCase interface {
	Execute(payload entity.UserIDPayload) (entity.User, int, error)
}

type GetUserDoctorByIDUseCase interface {
	Execute(payload entity.UserIDPayload) (entity.User, int, error)
}

type GetAuthenticatedUserUseCase interface {
	Execute(authorizationHeader entity.AuthorizationHeader) (entity.User, int, error)
}
