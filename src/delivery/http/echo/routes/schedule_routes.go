package routes

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/handler"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/middleware"
	repository "github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/repository/postgres"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/jwt"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/nanoid"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/postgres"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/usecase/schedule"

	"github.com/labstack/echo/v4"
)

func ScheduleRoutes(e *echo.Echo) {
	postgresDB := postgres.Connect()
	scheduleRepository := repository.NewScheduleRepository(postgresDB)
	userRepository := repository.NewUserRepository(postgresDB)

	jwtTokenManager := jwt.NewJWTTokenManager()
	nanoidIDGenerator := nanoid.NewNanoidIDGenerator()

	addScheduleUseCase := schedule.NewAddScheduleUseCase(
		scheduleRepository,
		userRepository,
		jwtTokenManager,
		nanoidIDGenerator,
	)
	getSchedulesUseCase := schedule.NewGetSchedulesUseCase(scheduleRepository)
	getScheduleByIDUseCase := schedule.NewGetScheduleByIDUseCase(
		scheduleRepository,
		userRepository,
		jwtTokenManager,
	)
	updateScheduleByIDUseCase := schedule.NewUpdateScheduleByIDUseCase(
		scheduleRepository,
		userRepository,
		jwtTokenManager,
	)
	deleteScheduleByIDUseCase := schedule.NewDeleteScheduleByIDUseCase(
		scheduleRepository,
		userRepository,
		jwtTokenManager,
	)
	getScheduleByDoctorIDUseCase := schedule.NewGetSchedulesByDoctorIDUseCase(
		scheduleRepository,
		userRepository,
		jwtTokenManager,
	)

	scheduleHandler := handler.NewScheduleHandler(
		addScheduleUseCase,
		getSchedulesUseCase,
		getScheduleByIDUseCase,
		updateScheduleByIDUseCase,
		deleteScheduleByIDUseCase,
		getScheduleByDoctorIDUseCase,
	)

	e.POST("/users/:userID/schedules", scheduleHandler.PostScheduleHandler, middleware.JWTMiddleware())
	e.GET("/users/:userID/schedules", scheduleHandler.GetSchedulesByDoctorIDHandler, middleware.JWTMiddleware())
	e.GET("/users/:userID/schedules/:scheduleID", scheduleHandler.GetScheduleByIDHandler, middleware.JWTMiddleware())
	e.PUT("/users/:userID/schedules/:scheduleID", scheduleHandler.PutScheduleByIDHandler, middleware.JWTMiddleware())
	e.DELETE("/users/:userID/schedules", scheduleHandler.DeleteScheduleByIDHandler, middleware.JWTMiddleware())
}
