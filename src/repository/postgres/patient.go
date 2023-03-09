package postgres

import (
	"fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"gorm.io/gorm"
)

type patientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) domain.PatientRepository {
	return &patientRepository{
		db: db,
	}
}

func (r *patientRepository) AddPatient(payload entity.Patient) (entity.AddedPatient, int, error) {
	result := r.db.Create(&payload)

	if result.RowsAffected == 0 {
		return entity.AddedPatient{}, http.StatusInternalServerError, result.Error
	}

	addedPatient := entity.AddedPatient{
		ID: payload.ID,
	}

	return addedPatient, http.StatusOK, nil
}

func (r *patientRepository) GetPatients() ([]entity.Patient, int, error) {
	Patients := []entity.Patient{}
	r.db.Find(&Patients)

	return Patients, http.StatusOK, nil
}

func (r *patientRepository) GetPatientByID(id string) (entity.Patient, int, error) {
	patient := entity.Patient{}
	result := r.db.Where("id = ?", id).First(&patient)

	if result.RowsAffected == 0 {
		return entity.Patient{}, http.StatusNotFound, fmt.Errorf("patient not found")
	}

	return patient, http.StatusOK, nil
}

func (r *patientRepository) UpdatePatientByID(payload entity.UpdatePatientPayload) (int, error) {
	patient, code, err := r.GetPatientByID(payload.ID)
	if err != nil {
		return code, err
	}

	patient.NIK = payload.NIK
	patient.Name = payload.Name
	patient.Phone = payload.Phone
	patient.Gender = payload.Gender
	patient.MedicalRecord = payload.MedicalRecord

	result := r.db.Save(&patient)

	if result.Error != nil {
		return http.StatusInternalServerError, result.Error
	}

	return http.StatusOK, nil
}

func (r *patientRepository) DeletePatientByID(id string) (int, error) {
	result := r.db.Where("id = ?", id).Delete(&entity.Patient{})

	if result.RowsAffected == 0 {
		return http.StatusNotFound, fmt.Errorf("patient not found")
	}

	return http.StatusOK, nil
}

func (r *patientRepository) GetPatientByMedicalRecord(medicalRecord string) (entity.Patient, int, error) {
	patient := entity.Patient{}
	result := r.db.Where("medical_record = ?", medicalRecord).First(&patient)

	if result.RowsAffected == 0 {
		return entity.Patient{}, http.StatusNotFound, fmt.Errorf("patient not found")
	}

	return patient, http.StatusOK, nil
}

func (r *patientRepository) GetPatientByNIK(nik string) (entity.Patient, int, error) {
	patient := entity.Patient{}
	result := r.db.Where("nik = ?", nik).First(&patient)

	if result.RowsAffected == 0 {
		return entity.Patient{}, http.StatusNotFound, fmt.Errorf("patient not found")
	}

	return patient, http.StatusOK, nil
}

func (r *patientRepository) VerifyNewNIK(nik string) (int, error) {
	result := r.db.Where("nik = ?", nik).First(&entity.Patient{})

	if result.RowsAffected > 0 {
		return http.StatusBadRequest, fmt.Errorf("nik already registered")
	}

	return http.StatusOK, nil
}
