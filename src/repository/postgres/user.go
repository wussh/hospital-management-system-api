package postgres

import (
	"fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) AddUser(payload entity.User) (entity.AddedUser, int, error) {
	result := r.db.Create(&payload)

	if result.RowsAffected == 0 {
		return entity.AddedUser{}, http.StatusInternalServerError, result.Error
	}

	addedUser := entity.AddedUser{
		ID:   payload.ID,
		Name: payload.Name,
	}

	return addedUser, http.StatusOK, nil
}

func (r *userRepository) GetUserByID(id string) (entity.User, int, error) {
	user := entity.User{}
	result := r.db.Where("id = ?", id).First(&user)

	if result.RowsAffected == 0 {
		return entity.User{}, http.StatusNotFound, fmt.Errorf("user not found")
	}

	return user, http.StatusOK, nil
}

func (r *userRepository) UpdateUserAvatar(payload entity.UpdateAvatarLocationPayload) (int, error) {
	user, code, err := r.GetUserByID(payload.ID)
	if err != nil {
		return code, err
	}

	user.Avatar = payload.Avatar
	result := r.db.Save(&user)

	if result.RowsAffected == 0 {
		return http.StatusInternalServerError, result.Error
	}

	return http.StatusOK, nil
}

func (r *userRepository) DeleteUserAvatar(payload entity.UserIDPayload) (int, error) {
	user, code, err := r.GetUserByID(payload.ID)
	if err != nil {
		return code, err
	}

	user.Avatar = ""
	result := r.db.Save(&user)

	if result.RowsAffected == 0 {
		return http.StatusInternalServerError, result.Error
	}

	return http.StatusOK, nil
}

func (r *userRepository) GetUserDoctorByID(id string) (entity.User, int, error) {
	user, code, err := r.GetUserByID(id)
	if err != nil {
		return entity.User{}, code, err
	}

	if user.Role != "Doctor" {
		return entity.User{}, http.StatusBadRequest, fmt.Errorf("user isn't Doctor")
	}

	return user, http.StatusOK, nil
}

func (r *userRepository) GetUserStaffByID(id string) (entity.User, int, error) {
	user, code, err := r.GetUserByID(id)
	if err != nil {
		return entity.User{}, code, err
	}

	if user.Role != "Staff" {
		return entity.User{}, http.StatusBadRequest, fmt.Errorf("user isn't Staff")
	}

	return user, http.StatusOK, nil
}

func (r *userRepository) GetUserDoctorsByClinicID(clinicID string) ([]entity.User, int, error) {
	users := []entity.User{}
	r.db.Where("role = ? AND clinic_id = ?", "Doctor", clinicID).Find(&users)

	return users, http.StatusOK, nil
}
