package postgres

import (
	"fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"gorm.io/gorm"
)

type clinicRepository struct {
	db *gorm.DB
}

func NewClinicRepository(db *gorm.DB) domain.ClinicRepository {
	return &clinicRepository{
		db: db,
	}
}

func (r *clinicRepository) AddClinic(payload entity.Clinic) (int, error) {
	result := r.db.Create(&payload)

	if result.RowsAffected == 0 {
		return http.StatusInternalServerError, result.Error
	}

	return http.StatusOK, nil
}

func (r *clinicRepository) GetClinics() ([]entity.Clinic, int, error) {
	clinics := []entity.Clinic{}
	r.db.Find(&clinics)

	return clinics, http.StatusOK, nil
}

func (r *clinicRepository) GetClinicByID(id string) (entity.Clinic, int, error) {
	clinic := entity.Clinic{}
	result := r.db.Where("id = ?", id).First(&clinic)

	if result.RowsAffected == 0 {
		return entity.Clinic{}, http.StatusNotFound, fmt.Errorf("clinic not found")
	}

	return clinic, http.StatusOK, nil
}

func (r *clinicRepository) UpdateClinicByID(payload entity.UpdateClinicPayload) (int, error) {
	Clinic, code, err := r.GetClinicByID(payload.ID)
	if err != nil {
		return code, err
	}

	Clinic.Name = payload.Name

	result := r.db.Save(&Clinic)

	if result.Error != nil {
		return http.StatusInternalServerError, result.Error
	}

	return http.StatusOK, nil
}

func (r *clinicRepository) DeleteClinicByID(id string) (int, error) {
	result := r.db.Where("id = ?", id).Delete(&entity.Clinic{})

	if result.RowsAffected == 0 {
		return http.StatusNotFound, fmt.Errorf("clinic not found")
	}

	return http.StatusOK, nil
}
