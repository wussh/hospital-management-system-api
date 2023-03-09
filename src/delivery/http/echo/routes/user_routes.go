package routes

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/handler"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/middleware"
	repository "github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/repository/postgres"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/bcrypt"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/jwt"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/nanoid"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/postgres"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/usecase/user"
	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	postgresDB := postgres.Connect()

	userRepository := repository.NewUserRepository(postgresDB)
	scheduleRepository := repository.NewScheduleRepository(postgresDB)
	dayRepository := repository.NewDayRepository(postgresDB)
	timeRepository := repository.NewTimeRepository(postgresDB)
	clinicRepository := repository.NewClinicRepository(postgresDB)

	jwtTokenManager := jwt.NewJWTTokenManager()
	bcryptPasswordHash := bcrypt.NewBcryptPasswordHash()
	nanoidIDGenerator := nanoid.NewNanoidIDGenerator()

	addUserUseCase := user.NewAddUserUseCase(
		userRepository,
		bcryptPasswordHash,
		jwtTokenManager,
		nanoidIDGenerator,
	)
	updateUserAvatarUseCase := user.NewUpdateUserAvatarUseCase(userRepository, jwtTokenManager)
	deleteUserAvatarUseCase := user.NewDeleteUserAvatarUseCase(userRepository, jwtTokenManager)
	getUserByIDUseCase := user.NewGetUserByIDUseCase(userRepository)
	getUserDoctorByIDUseCase := user.NewGetUserDoctorByIDUseCase(
		userRepository,
		scheduleRepository,
		dayRepository,
		timeRepository,
		clinicRepository,
	)
	getAuthenticatedUserUseCase := user.NewGetAuthenticatedUserUseCase(
		userRepository,
		jwtTokenManager,
	)

	userHandler := handler.NewUserHandler(
		addUserUseCase,
		updateUserAvatarUseCase,
		deleteUserAvatarUseCase,
		getUserByIDUseCase,
		getUserDoctorByIDUseCase,
		getAuthenticatedUserUseCase,
	)

	e.POST("/users", userHandler.PostUserHandler, middleware.JWTMiddleware())
	e.PUT("/users/:userID/avatar", userHandler.PutUserAvatarHandler, middleware.JWTMiddleware())
	e.DELETE("/users/:userID/avatar", userHandler.DeleteUserAvatarHandler, middleware.JWTMiddleware())
	e.GET("/users/:userID", userHandler.GetUserByIDHandler, middleware.JWTMiddleware())
	e.GET("/users/current", userHandler.GetAuthenticatedUserHandler, middleware.JWTMiddleware())
}
