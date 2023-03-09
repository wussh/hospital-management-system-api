package patient

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
)

type deletePatientByIDUseCase struct {
	patientRepository domain.PatientRepository
}

func NewDeletePatientByIDUseCase(patientRepository domain.PatientRepository) domain.DeletePatientByIDUseCase {
	return &deletePatientByIDUseCase{
		patientRepository: patientRepository,
	}
}

func (u *deletePatientByIDUseCase) Execute(id string) (int, error) {
	return u.patientRepository.DeletePatientByID(id)
}
