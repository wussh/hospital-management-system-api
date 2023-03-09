package doctor

import (
	// "fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
)

type addDoctorUseCase struct {
	doctorRepository domain.DoctorRepository
	clinicRepository domain.ClinicRepository
	jwtTokenManager  application.TokenManager
	staffRepository  domain.StaffRepository
}

func NewAddDoctorUseCase(
	doctorRepository domain.DoctorRepository,
	clinicRepository domain.ClinicRepository,
	jwtTokenManager application.TokenManager,
	staffRepository domain.StaffRepository,
) domain.AddDoctorUseCase {
	return &addDoctorUseCase{
		doctorRepository: doctorRepository,
		clinicRepository: clinicRepository,
		jwtTokenManager:  jwtTokenManager,
		staffRepository:  staffRepository,
	}
}

func (u *addDoctorUseCase) Execute(
	payload entity.Doctor,
	authorizationHeader entity.AuthorizationHeader,
) (int, error) {
	return http.StatusOK, nil

	// decodedPayload, code, err := u.jwtTokenManager.DecodeAccessTokenPayload(authorizationHeader.AccessToken)
	// if err != nil {
	// 	return code, err
	// }

	// staff, code, err := u.staffRepository.GetStaffByID(decodedPayload.ID)
	// if err != nil {
	// 	return code, err
	// }

	// if staff.StaffType != "admin" {
	// 	return http.StatusForbidden, fmt.Errorf("restricted resource")
	// }

	// if _, code, err := u.clinicRepository.GetClinicByID(payload.ClinicID); err != nil {
	// 	return code, err
	// }

	return u.doctorRepository.AddDoctor(payload)
}
