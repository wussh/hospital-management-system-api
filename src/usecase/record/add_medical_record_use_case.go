package record

import (
	"fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
)

type addMedicalRecordUseCase struct {
	medicalRecordRepository domain.MedicalRecordRepository
	nanoidIDGenerator       application.IDGenerator
	patientRepository       domain.PatientRepository
	sessionRepository       domain.SessionRepository
}

func NewAddMedicalRecordUseCase(
	medicalRecordRepository domain.MedicalRecordRepository,
	nanoidIDGenerator application.IDGenerator,
	patientRepository domain.PatientRepository,
	sessionRepository domain.SessionRepository,
) domain.AddMedicalRecordUseCase {
	return &addMedicalRecordUseCase{
		medicalRecordRepository: medicalRecordRepository,
		nanoidIDGenerator:       nanoidIDGenerator,
		patientRepository:       patientRepository,
		sessionRepository:       sessionRepository,
	}
}

func (u *addMedicalRecordUseCase) Execute(payload entity.MedicalRecord) (int, error) {
	if payload.Type != "Rawat Jalan" {
		return http.StatusBadRequest, fmt.Errorf("type must be Rawat Jalan")
	}

	if payload.Status != "Rawat Jalan" && payload.Status != "Pulang" {
		return http.StatusBadRequest, fmt.Errorf("status must be Rawat Jalan or Pulang")
	}

	if _, code, err := u.patientRepository.GetPatientByMedicalRecord(*payload.PatientMedicalRecord); err != nil {
		return code, err
	}

	if _, code, err := u.sessionRepository.GetSessionByID(*payload.SessionID); err != nil {
		return code, err
	}

	generatedID, code, err := u.nanoidIDGenerator.Generate()
	if err != nil {
		return code, err
	}
	payload.ID = fmt.Sprintf("medical-record-%s", generatedID)

	return u.medicalRecordRepository.AddMedicalRecord(payload)
}
