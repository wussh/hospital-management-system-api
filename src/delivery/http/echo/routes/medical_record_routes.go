package routes

import (
	// "github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/handler"
	// "github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/middleware"
	// repository "github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/repository/postgres"
	// "github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/nanoid"
	// "github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/postgres"
	// "github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/usecase/record"
	"github.com/labstack/echo/v4"
)

func MedicalRecordRoutes(e *echo.Echo) {
	// postgresDB := postgres.Connect()

	// medicalRecordRepository := repository.NewMedicalRecordRepository(postgresDB)
	// patientRepository := repository.NewPatientRepository(postgresDB)
	// sessionRepository := repository.NewSessionRepository(postgresDB)

	// nanoidIDGenerator := nanoid.NewNanoidIDGenerator()

	// addMedicalRecordUseCase := record.NewAddMedicalRecordUseCase(
	// 	medicalRecordRepository,
	// 	nanoidIDGenerator,
	// 	patientRepository,
	// 	sessionRepository,
	// )

	// patientHandler := handler.NewMedicalRecordHandler(
	// 	addMedicalRecordUseCase,
	// )

	// e.POST("/sessions/:sessionID/records", patientHandler.PostMedicalRecordHandler, middleware.JWTMiddleware())
}
