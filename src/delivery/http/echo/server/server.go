package server

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/routes"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CreateServer() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORS())

	routes.UserRoutes(e)
	routes.AuthenticationRoutes(e)
	routes.ClinicRoutes(e)
	routes.ScheduleRoutes(e)
	routes.PatientRoutes(e)
	routes.SessionRoutes(e)
	routes.MedicalRecordRoutes(e)

	// routes.HelloRoutes(e)
	// routes.StaffRoutes(e)
	// routes.DoctorRoutes(e)

	routes.StaticFiles(e)

	validator.NewValidator(e)

	return e
}
