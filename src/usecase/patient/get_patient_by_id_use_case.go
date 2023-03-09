package patient

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type getPatientByIDUseCase struct {
	patientRepository       domain.PatientRepository
	sessionRepository       domain.SessionRepository
	medicalRecordRepository domain.MedicalRecordRepository
}

func NewGetPatientByIDUseCase(
	patientRepository domain.PatientRepository,
	sessionRepository domain.SessionRepository,
	medicalRecordRepository domain.MedicalRecordRepository,
) domain.GetPatientByIDUseCase {
	return &getPatientByIDUseCase{
		patientRepository:       patientRepository,
		sessionRepository:       sessionRepository,
		medicalRecordRepository: medicalRecordRepository,
	}
}

func (u *getPatientByIDUseCase) Execute(id string) (entity.Patient, int, error) {
	patient, code, err := u.patientRepository.GetPatientByID(id)
	if err != nil {
		return entity.Patient{}, code, err
	}

	sessions, code, err := u.sessionRepository.GetSessionByPatientID(id)
	if err != nil {
		return entity.Patient{}, code, err
	}
	patient.Sessions = sessions

	records, code, err := u.medicalRecordRepository.GetMedicalRecordsByPatientMedicalRecord(patient.MedicalRecord)
	if err != nil {
		return entity.Patient{}, code, err
	}
	patient.MedicalRecords = records

	return patient, code, err
}
