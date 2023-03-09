package staff

import (
	"fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
)

type addStaffUseCase struct {
	staffRepository    domain.StaffRepository
	bcryptPasswordHash application.PasswordHash
	jwtTokenManager    application.TokenManager
}

func NewAddStaffUseCase(
	staffRepository domain.StaffRepository,
	bcryptPasswordHash application.PasswordHash,
	jwtTokenManager application.TokenManager,
) domain.AddStaffUseCase {
	return &addStaffUseCase{
		staffRepository:    staffRepository,
		bcryptPasswordHash: bcryptPasswordHash,
		jwtTokenManager:    jwtTokenManager,
	}
}

func (u *addStaffUseCase) Execute(
	payload entity.Staff,
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

	if payload.StaffType == "admin" {
		return http.StatusBadRequest, fmt.Errorf("admin staff already exists")
	}

	if code, err := u.staffRepository.VerifyEmailAvailable(payload.Email); err != nil {
		return code, err
	}

	hashedPassword, code, err := u.bcryptPasswordHash.Hash(payload.Password)
	if err != nil {
		return code, err
	}

	payload.Password = hashedPassword

	return u.staffRepository.AddStaff(payload)
}
