package doctor

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type getDoctorsUseCase struct {
	doctorRepository domain.DoctorRepository
}

func NewGetDoctorsUseCase(doctorRepository domain.DoctorRepository) domain.GetDoctorsUseCase {
	return &getDoctorsUseCase{
		doctorRepository: doctorRepository,
	}
}

func (u *getDoctorsUseCase) Execute() ([]entity.Doctor, int, error) {
	return u.doctorRepository.GetDoctors()
}
