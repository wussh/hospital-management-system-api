package routes

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/handler"
	repository "github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/repository/postgres"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/bcrypt"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/jwt"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/postgres"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/usecase/staff"
	"github.com/labstack/echo/v4"
)

func StaffRoutes(e *echo.Echo) {
	postgresDB := postgres.Connect()

	staffRepository := repository.NewStaffRepository(postgresDB)
	jwtTokenManager := jwt.NewJWTTokenManager()

	bcryptPasswordHash := bcrypt.NewBcryptPasswordHash()

	addStaffUseCase := staff.NewAddStaffUseCase(staffRepository, bcryptPasswordHash, jwtTokenManager)

	staffHandler := handler.NewStaffHandler(addStaffUseCase)

	e.POST("/staffs", staffHandler.PostStaffHandler)
}
