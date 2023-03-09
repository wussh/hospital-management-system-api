package patient

import (
	"fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
)

type addPatientUseCase struct {
	patientRepository domain.PatientRepository
	nanoidIDGenerator application.IDGenerator
}

func NewAddPatientUseCase(
	patientRepository domain.PatientRepository,
	nanoidIDGenerator application.IDGenerator,
) domain.AddPatientUseCase {
	return &addPatientUseCase{
		patientRepository: patientRepository,
		nanoidIDGenerator: nanoidIDGenerator,
	}
}

func (u *addPatientUseCase) Execute(payload entity.Patient) (entity.AddedPatient, int, error) {
	if payload.Gender != "Laki-Laki" && payload.Gender != "Perempuan" {
		return entity.AddedPatient{}, http.StatusBadRequest, fmt.Errorf("gender must be Laki-Laki or Perempuan")
	}

	if code, err := u.patientRepository.VerifyNewNIK(payload.NIK); err != nil {
		return entity.AddedPatient{}, code, err
	}

	generatedID, code, err := u.nanoidIDGenerator.Generate()
	if err != nil {
		return entity.AddedPatient{}, code, err
	}
	payload.ID = fmt.Sprintf("patient-%s", generatedID)

	payload.MedicalRecord = fmt.Sprintf("record-%s-%s", generatedID, payload.NIK)

	return u.patientRepository.AddPatient(payload)
}
