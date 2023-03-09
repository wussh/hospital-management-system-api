package patient

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type getPatientsUseCase struct {
	patientRepository domain.PatientRepository
}

func NewGetPatientsUseCase(patientRepository domain.PatientRepository) domain.GetPatientsUseCase {
	return &getPatientsUseCase{
		patientRepository: patientRepository,
	}
}

func (u *getPatientsUseCase) Execute() ([]entity.Patient, int, error) {
	return u.patientRepository.GetPatients()
}
