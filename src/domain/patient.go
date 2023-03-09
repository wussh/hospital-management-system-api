package domain

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/labstack/echo/v4"
)

type PatientHandler interface {
	PostPatientHandler(c echo.Context) error
	GetPatientsHandler(c echo.Context) error
	GetPatientByIDHandler(c echo.Context) error
	PutPatientByIDHandler(c echo.Context) error
	DeletePatientByIDHandler(c echo.Context) error
	GetPatientByNIKHandler(c echo.Context) error
}

type PatientRepository interface {
	AddPatient(payload entity.Patient) (entity.AddedPatient, int, error)
	GetPatients() ([]entity.Patient, int, error)
	GetPatientByID(id string) (entity.Patient, int, error)
	UpdatePatientByID(payload entity.UpdatePatientPayload) (int, error)
	DeletePatientByID(id string) (int, error)
	GetPatientByMedicalRecord(medicalRecord string) (entity.Patient, int, error)
	GetPatientByNIK(nik string) (entity.Patient, int, error)
	VerifyNewNIK(nik string) (int, error)
}

type AddPatientUseCase interface {
	Execute(payload entity.Patient) (entity.AddedPatient, int, error)
}

type GetPatientsUseCase interface {
	Execute() ([]entity.Patient, int, error)
}

type GetPatientByIDUseCase interface {
	Execute(id string) (entity.Patient, int, error)
}

type UpdatePatientByIDUseCase interface {
	Execute(payload entity.UpdatePatientPayload) (int, error)
}

type DeletePatientByIDUseCase interface {
	Execute(id string) (int, error)
}

type GetPatientByNIKUseCase interface {
	Execute(payload entity.PatientNIKPayload) (entity.Patient, int, error)
}
