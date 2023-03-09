package domain

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/labstack/echo/v4"
)

type StaffHandler interface {
	PostStaffHandler(c echo.Context) error
}

type StaffRepository interface {
	AddStaff(payload entity.Staff) (int, error)
	GetStaffByEmail(email string) (entity.Staff, int, error)
	VerifyEmailAvailable(email string) (int, error)
	GetStaffByID(id uint) (entity.Staff, int, error)
}

type AddStaffUseCase interface {
	Execute(payload entity.Staff, authorizationHeader entity.AuthorizationHeader) (int, error)
}
