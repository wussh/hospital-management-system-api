package clinic

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type getClinicsUseCase struct {
	clinicRepository domain.ClinicRepository
}

func NewGetClinicsUseCase(clinicRepository domain.ClinicRepository) domain.GetClinicsUseCase {
	return &getClinicsUseCase{
		clinicRepository: clinicRepository,
	}
}

func (u *getClinicsUseCase) Execute() ([]entity.Clinic, int, error) {
	return u.clinicRepository.GetClinics()
}
