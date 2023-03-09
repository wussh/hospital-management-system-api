package routes

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/handler"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/middleware"
	repository "github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/repository/postgres"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/jwt"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/nanoid"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/postgres"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/usecase/clinic"
	"github.com/labstack/echo/v4"
)

func ClinicRoutes(e *echo.Echo) {
	postgresDB := postgres.Connect()

	clinicRepository := repository.NewClinicRepository(postgresDB)
	userRepository := repository.NewUserRepository(postgresDB)

	jwtTokenManager := jwt.NewJWTTokenManager()
	nanoidIDGenerator := nanoid.NewNanoidIDGenerator()

	addClinicUseCase := clinic.NewAddClinicUseCase(
		clinicRepository,
		jwtTokenManager,
		userRepository,
		nanoidIDGenerator,
	)
	getClinicsUseCase := clinic.NewGetClinicsUseCase(clinicRepository)
	getClinicByIDUseCase := clinic.NewGetClinicByIDUseCase(clinicRepository, userRepository)
	updateClinicByIDUseCase := clinic.NewUpdateClinicByIDUseCase(
		clinicRepository,
		jwtTokenManager,
		userRepository,
	)
	deleteClinicByIDUseCase := clinic.NewDeleteClinicByIDUseCase(
		clinicRepository,
		jwtTokenManager,
		userRepository,
	)

	clinicHandler := handler.NewClinicHandler(
		addClinicUseCase,
		getClinicsUseCase,
		getClinicByIDUseCase,
		updateClinicByIDUseCase,
		deleteClinicByIDUseCase,
	)

	e.POST("/clinics", clinicHandler.PostClinicHandler, middleware.JWTMiddleware())
	e.GET("/clinics", clinicHandler.GetClinicsHandler, middleware.JWTMiddleware())
	e.GET("/clinics/:clinicID", clinicHandler.GetClinicByIDHandler, middleware.JWTMiddleware())
	e.PUT("/clinics/:clinicID", clinicHandler.PutClinicByIDHandler, middleware.JWTMiddleware())
	e.DELETE("/clinics/:clinicID", clinicHandler.DeleteClinicByIDHandler, middleware.JWTMiddleware())
}
