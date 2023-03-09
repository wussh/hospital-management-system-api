package patient

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type updatePatientByIDUseCase struct {
	patientRepository domain.PatientRepository
}

func NewUpdatePatientByIDUseCase(patientRepository domain.PatientRepository) domain.UpdatePatientByIDUseCase {
	return &updatePatientByIDUseCase{
		patientRepository: patientRepository,
	}
}

func (u *updatePatientByIDUseCase) Execute(payload entity.UpdatePatientPayload) (int, error) {
	return u.patientRepository.UpdatePatientByID(payload)
}
