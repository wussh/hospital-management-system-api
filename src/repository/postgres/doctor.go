package postgres

import (
	"fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"gorm.io/gorm"
)

type doctorRepository struct {
	db *gorm.DB
}

func NewDoctorRepository(db *gorm.DB) domain.DoctorRepository {
	return &doctorRepository{
		db: db,
	}
}

func (r *doctorRepository) AddDoctor(payload entity.Doctor) (int, error) {
	result := r.db.Create(&payload)

	if result.RowsAffected == 0 {
		return http.StatusInternalServerError, result.Error
	}

	return http.StatusOK, nil
}

func (r *doctorRepository) GetDoctors() ([]entity.Doctor, int, error) {
	doctors := []entity.Doctor{}
	r.db.Find(&doctors)

	return doctors, http.StatusOK, nil
}

func (r *doctorRepository) GetDoctorByID(id uint) (entity.Doctor, int, error) {
	doctor := entity.Doctor{}
	result := r.db.Where("id = ?", id).First(&doctor)

	if result.RowsAffected == 0 {
		return entity.Doctor{}, http.StatusNotFound, fmt.Errorf("doctor not found")
	}

	return doctor, http.StatusOK, nil
}

func (r *doctorRepository) UpdateDoctorByID(payload entity.UpdateDoctorPayload) (int, error) {
	doctor, code, err := r.GetDoctorByID(payload.ID)
	if err != nil {
		return code, err
	}

	doctor.FName = payload.FName
	doctor.LName = payload.LName
	doctor.Phone = payload.Phone
	doctor.ClinicID = payload.ClinicID

	result := r.db.Save(&doctor)

	if result.Error != nil {
		return http.StatusInternalServerError, result.Error
	}

	return http.StatusOK, nil
}

func (r *doctorRepository) DeleteDoctorByID(id uint) (int, error) {
	result := r.db.Where("id = ?", id).Delete(&entity.Doctor{})

	if result.RowsAffected == 0 {
		return http.StatusNotFound, fmt.Errorf("doctor not found")
	}

	return http.StatusOK, nil
}
