package staff_test

import (
	"testing"
	// "fmt"
	// "net/http"
	// domainmock "github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain/mocks"
	// "github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	// applicationmock "github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application/mocks"
	// "github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/usecase/staff"
	// "github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/mock"
)

func TestAddStaffUseCase(t *testing.T) {
	// mockStaffRepository := &domainmock.StaffRepository{}
	// mockPasswordHash := &applicationmock.PasswordHash{}

	// t.Run("throw error email not available", func(t *testing.T) {
	// 	useCasePayload := entity.Staff{
	// 		StaffType: "staff",
	// 		Name:      "Staff",
	// 		Phone:     "08123456789",
	// 		Email:     "unavailable@simars.com",
	// 		Password:  "plain_password",
	// 		Token:     "staff_token",
	// 	}

	// 	mockStaffRepository.On("VerifyEmailAvailable", useCasePayload.Email).
	// 		Return(http.StatusBadRequest, fmt.Errorf("email already registered")).Once()
	// 	mockPasswordHash.On("Hash", mock.Anything).Return(mock.Anything, mock.Anything).Once()
	// 	mockStaffRepository.On("AddStaff", mock.Anything).Return(mock.Anything, mock.Anything).Once()

	// 	addStaffUseCase := staff.NewAddStaffUseCase(mockStaffRepository, mockPasswordHash)
	// 	code, err := addStaffUseCase.Execute(useCasePayload)

	// 	assert := assert.New(t)
	// 	assert.Equal(http.StatusBadRequest, code)
	// 	assert.Error(err)
	// 	assert.Equal(fmt.Errorf("email already registered"), err)
	// })

	// t.Run("success", func(t *testing.T) {
	// 	useCasePayload := entity.Staff{
	// 		StaffType: "staff",
	// 		Name:      "Staff",
	// 		Phone:     "08123456789",
	// 		Email:     "staff@simars.com",
	// 		Password:  "plain_password",
	// 		Token:     "staff_token",
	// 	}

	// 	mockStaffRepository.On("VerifyEmailAvailable", useCasePayload.Email).Return(http.StatusOK, nil).Once()
	// 	mockPasswordHash.On("Hash", useCasePayload.Password).Return("encrypted_password", nil).Once()
	// 	mockStaffRepository.On("AddStaff", entity.Staff{
	// 		StaffType: "staff",
	// 		Name:      "Staff",
	// 		Phone:     "08123456789",
	// 		Email:     "staff@simars.com",
	// 		Password:  "encrypted_password",
	// 		Token:     "staff_token",
	// 	}).Return(http.StatusOK, nil).Once()

	// 	addStaffUseCase := staff.NewAddStaffUseCase(mockStaffRepository, mockPasswordHash)
	// 	code, err := addStaffUseCase.Execute(useCasePayload)

	// 	assert := assert.New(t)
	// 	assert.Equal(http.StatusOK, code)
	// 	assert.NoError(err)
	// })
}
