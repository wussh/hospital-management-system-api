package entity

import "time"

type MedicalRecord struct {
	ID                   string    `gorm:"not null" json:"id"`
	SessionID            *string   `gorm:"not null" param:"sessionID" json:"sessionID" validate:"required"`
	PatientMedicalRecord *string   `gorm:"not null" json:"patientMedicalRecord"`
	Type                 string    `gorm:"not null" json:"type" validate:"required"`
	History              string    `gorm:"not null" json:"history" validate:"required"`
	Diagnosis            string    `gorm:"not null" json:"diagnosis" validate:"required"`
	DrugAllergyHistory   string    `gorm:"not null" json:"drugAllergyHistory" validate:"required"`
	DrugTherapy          string    `gorm:"not null" json:"drugTherapy" validate:"required"`
	Height               string    `gorm:"not null" json:"height" validate:"required,number,gt=0"`
	Weight               string    `gorm:"not null" json:"weight" validate:"required,number,gt=0"`
	Systole              string    `gorm:"not null" json:"systole" validate:"required,number,gt=0"`
	Diastole             string    `gorm:"not null" json:"diastole" validate:"required,number,gt=0"`
	Temperature          string    `gorm:"not null" json:"temperature" validate:"required,number,gt=0"`
	Status               string    `gorm:"not null" json:"status" validate:"required"`
	CreatedAt            time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt            time.Time `gorm:"not null" json:"updatedAt"`
}
