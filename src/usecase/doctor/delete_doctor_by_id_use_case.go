package doctor

import (
	// "fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
)

type deleteDoctorByIDUseCase struct {
	doctorRepository domain.DoctorRepository
	jwtTokenManager  application.TokenManager
	staffRepository  domain.StaffRepository
}

func NewDeleteDoctorByIDUseCase(
	doctorRepository domain.DoctorRepository,
	jwtTokenManager application.TokenManager,
	staffRepository domain.StaffRepository,
) domain.DeleteDoctorByIDUseCase {
	return &deleteDoctorByIDUseCase{
		doctorRepository: doctorRepository,
		jwtTokenManager:  jwtTokenManager,
		staffRepository:  staffRepository,
	}
}

func (u *deleteDoctorByIDUseCase) Execute(
	id uint,
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
	return u.doctorRepository.DeleteDoctorByID(id)
}
