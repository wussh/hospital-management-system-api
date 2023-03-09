package patient

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type getPatientByNIKUseCase struct {
	patientRepository domain.PatientRepository
}

func NewGetPatientByNIKUseCase(
	patientRepository domain.PatientRepository,
) domain.GetPatientByNIKUseCase {
	return &getPatientByNIKUseCase{
		patientRepository: patientRepository,
	}
}

func (u *getPatientByNIKUseCase) Execute(payload entity.PatientNIKPayload) (entity.Patient, int, error) {
	patient, code, err := u.patientRepository.GetPatientByNIK(payload.NIK)
	if err != nil {
		return entity.Patient{}, code, err
	}

	return patient, code, err
}
