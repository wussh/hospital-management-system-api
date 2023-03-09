package postgres

import (
	"fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"

	"gorm.io/gorm"
)

type staffRepository struct {
	db *gorm.DB
}

func NewStaffRepository(db *gorm.DB) domain.StaffRepository {
	return &staffRepository{
		db: db,
	}
}

func (r *staffRepository) AddStaff(payload entity.Staff) (int, error) {
	result := r.db.Create(&payload)

	if result.RowsAffected == 0 {
		return http.StatusInternalServerError, result.Error
	}

	return http.StatusOK, nil
}

func (r *staffRepository) GetStaffByEmail(email string) (entity.Staff, int, error) {
	staff := entity.Staff{}
	result := r.db.Where("email = ?", email).First(&staff)

	if result.RowsAffected == 0 {
		return entity.Staff{}, http.StatusNotFound, fmt.Errorf("staff not found")
	}

	return staff, http.StatusOK, nil
}

func (r *staffRepository) VerifyEmailAvailable(email string) (int, error) {
	staff := entity.Staff{}
	result := r.db.Where("email = ?", email).First(&staff)

	if result.RowsAffected > 0 {
		return http.StatusBadRequest, fmt.Errorf("email already registered")
	}

	return http.StatusOK, nil
}

func (r *staffRepository) GetStaffByID(id uint) (entity.Staff, int, error) {
	staff := entity.Staff{}
	result := r.db.Where("id = ?", id).First(&staff)

	if result.RowsAffected == 0 {
		return entity.Staff{}, http.StatusNotFound, fmt.Errorf("staff not found")
	}

	return staff, http.StatusOK, nil
}
