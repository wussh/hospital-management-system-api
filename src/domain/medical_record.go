package domain

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/labstack/echo/v4"
)

type MedicalRecordHandler interface {
	PostMedicalRecordHandler(c echo.Context) error
}

type MedicalRecordRepository interface {
	AddMedicalRecord(payload entity.MedicalRecord) (int, error)
	GetMedicalRecordsByPatientMedicalRecord(
		patientMedicalRecord string,
	) ([]entity.MedicalRecord, int, error)
}

type AddMedicalRecordUseCase interface {
	Execute(payload entity.MedicalRecord) (int, error)
}
