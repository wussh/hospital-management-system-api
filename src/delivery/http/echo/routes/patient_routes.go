package routes

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/handler"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/middleware"
	repository "github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/repository/postgres"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/nanoid"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/postgres"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/usecase/patient"

	"github.com/labstack/echo/v4"
)

func PatientRoutes(e *echo.Echo) {
	postgresDB := postgres.Connect()

	patientRepository := repository.NewPatientRepository(postgresDB)
	sessionRepository := repository.NewSessionRepository(postgresDB)
	medicalRecordRepository := repository.NewMedicalRecordRepository(postgresDB)

	nanoidIDGenerator := nanoid.NewNanoidIDGenerator()

	addPatientUseCase := patient.NewAddPatientUseCase(patientRepository, nanoidIDGenerator)
	getPatientsUseCase := patient.NewGetPatientsUseCase(patientRepository)
	getPatientByIDUseCase := patient.NewGetPatientByIDUseCase(
		patientRepository,
		sessionRepository,
		medicalRecordRepository,
	)
	updatePatientByIDUseCase := patient.NewUpdatePatientByIDUseCase(patientRepository)
	deletePatientByIDUseCase := patient.NewDeletePatientByIDUseCase(patientRepository)
	getPatientByNIKUseCase := patient.NewGetPatientByNIKUseCase(patientRepository)

	patientHandler := handler.NewPatientHandler(
		addPatientUseCase,
		getPatientsUseCase,
		getPatientByIDUseCase,
		updatePatientByIDUseCase,
		deletePatientByIDUseCase,
		getPatientByNIKUseCase,
	)

	e.POST("/patients", patientHandler.PostPatientHandler, middleware.JWTMiddleware())
	e.GET("/patients", patientHandler.GetPatientsHandler, middleware.JWTMiddleware())
	e.GET("/patients/:patientID", patientHandler.GetPatientByIDHandler, middleware.JWTMiddleware())
	e.PUT("/patients/:patientID", patientHandler.PutPatientByIDHandler, middleware.JWTMiddleware())
	e.DELETE("/patients/:patientID", patientHandler.DeletePatientByIDHandler, middleware.JWTMiddleware())
	e.GET("/nik/:nik", patientHandler.GetPatientByNIKHandler, middleware.JWTMiddleware())
}
