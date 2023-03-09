package postgres

import (
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"gorm.io/gorm"
)

type recordRepository struct {
	db *gorm.DB
}

func NewMedicalRecordRepository(db *gorm.DB) domain.MedicalRecordRepository {
	return &recordRepository{
		db: db,
	}
}

func (r *recordRepository) AddMedicalRecord(payload entity.MedicalRecord) (int, error) {
	result := r.db.Create(&payload)

	if result.RowsAffected == 0 {
		return http.StatusInternalServerError, result.Error
	}

	return http.StatusOK, nil
}

func (r *recordRepository) GetMedicalRecordsByPatientMedicalRecord(
	patientMedicalRecord string,
) ([]entity.MedicalRecord, int, error) {
	records := []entity.MedicalRecord{}
	r.db.Where("patient_medical_record = ?", patientMedicalRecord).Find(&records)

	return records, http.StatusOK, nil
}
