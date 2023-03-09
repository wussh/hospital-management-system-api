package domain

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/labstack/echo/v4"
)

type DoctorHandler interface {
	PostDoctorHandler(c echo.Context) error
	GetDoctorsHandler(c echo.Context) error
	GetDoctorByIDHandler(c echo.Context) error
	PutDoctorByIDHandler(c echo.Context) error
	DeleteDoctorByIDHandler(c echo.Context) error
}

type DoctorRepository interface {
	AddDoctor(payload entity.Doctor) (int, error)
	GetDoctors() ([]entity.Doctor, int, error)
	GetDoctorByID(id uint) (entity.Doctor, int, error)
	UpdateDoctorByID(payload entity.UpdateDoctorPayload) (int, error)
	DeleteDoctorByID(id uint) (int, error)
}

type AddDoctorUseCase interface {
	Execute(payload entity.Doctor, authorizationHeader entity.AuthorizationHeader) (int, error)
}

type GetDoctorsUseCase interface {
	Execute() ([]entity.Doctor, int, error)
}

type GetDoctorByIDUseCase interface {
	Execute(id uint) (entity.Doctor, int, error)
}

type UpdateDoctorByIDUseCase interface {
	Execute(
		payload entity.UpdateDoctorPayload,
		authorizationHeader entity.AuthorizationHeader,
	) (int, error)
}

type DeleteDoctorByIDUseCase interface {
	Execute(id uint, authorizationHeader entity.AuthorizationHeader) (int, error)
}
