package domain

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"

	"github.com/labstack/echo/v4"
)

type AuthenticationHandler interface {
	PostStaffLoginHandler(c echo.Context) error
	PostStaffLogoutHandler(c echo.Context) error
	PutAuthenticationHandler(c echo.Context) error
}

type AuthenticationRepository interface {
	AddRefreshToken(payload entity.Authentication) (int, error)
	VerifyRefreshTokenExistence(payload entity.RefreshTokenPayload) (int, error)
	DeleteRefreshToken(payload entity.RefreshTokenPayload) (int, error)
}

type StaffLoginUseCase interface {
	Execute(payload entity.LoginPayload) (entity.NewLogin, int, error)
}

type StaffLogoutUseCase interface {
	Execute(payload entity.RefreshTokenPayload) (int, error)
}

type UpdateAuthenticationUseCase interface {
	Execute(payload entity.RefreshTokenPayload) (entity.NewAccessToken, int, error)
}
