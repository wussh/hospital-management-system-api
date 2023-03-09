package clinic

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type getClinicByIDUseCase struct {
	clinicRepository domain.ClinicRepository
	userRepository   domain.UserRepository
}

func NewGetClinicByIDUseCase(
	clinicRepository domain.ClinicRepository,
	userRepository domain.UserRepository,
) domain.GetClinicByIDUseCase {
	return &getClinicByIDUseCase{
		clinicRepository: clinicRepository,
		userRepository:   userRepository,
	}
}

func (u *getClinicByIDUseCase) Execute(id string) (entity.Clinic, int, error) {
	clinic, code, err := u.clinicRepository.GetClinicByID(id)
	if err != nil {
		return entity.Clinic{}, code, err
	}

	doctors, code, err := u.userRepository.GetUserDoctorsByClinicID(id)
	if err != nil {
		return entity.Clinic{}, code, err
	}
	clinic.Doctors = doctors

	return clinic, code, err
}
