package domain

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/labstack/echo/v4"
)

type ClinicHandler interface {
	PostClinicHandler(c echo.Context) error
	GetClinicsHandler(c echo.Context) error
	GetClinicByIDHandler(c echo.Context) error
	PutClinicByIDHandler(c echo.Context) error
	DeleteClinicByIDHandler(c echo.Context) error
}

type ClinicRepository interface {
	AddClinic(payload entity.Clinic) (int, error)
	GetClinics() ([]entity.Clinic, int, error)
	GetClinicByID(id string) (entity.Clinic, int, error)
	UpdateClinicByID(payload entity.UpdateClinicPayload) (int, error)
	DeleteClinicByID(id string) (int, error)
}

type AddClinicUseCase interface {
	Execute(
		payload entity.Clinic,
		authorizationHeader entity.AuthorizationHeader,
	) (int, error)
}

type GetClinicsUseCase interface {
	Execute() ([]entity.Clinic, int, error)
}

type GetClinicByIDUseCase interface {
	Execute(id string) (entity.Clinic, int, error)
}

type UpdateClinicByIDUseCase interface {
	Execute(
		payload entity.UpdateClinicPayload,
		authorizationHeader entity.AuthorizationHeader,
	) (int, error)
}

type DeleteClinicByIDUseCase interface {
	Execute(id string, authorizationHeader entity.AuthorizationHeader) (int, error)
}
