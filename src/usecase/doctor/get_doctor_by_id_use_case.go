package doctor

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type getDoctorByIDUseCase struct {
	doctorRepository domain.DoctorRepository
}

func NewGetDoctorByIDUseCase(doctorRepository domain.DoctorRepository) domain.GetDoctorByIDUseCase {
	return &getDoctorByIDUseCase{
		doctorRepository: doctorRepository,
	}
}

func (u *getDoctorByIDUseCase) Execute(id uint) (entity.Doctor, int, error) {
	return u.doctorRepository.GetDoctorByID(id)
}
