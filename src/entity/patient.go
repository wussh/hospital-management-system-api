package entity

import "time"

type Patient struct {
	ID             string          `gorm:"not null" json:"id"`
	NIK            string          `gorm:"not null" json:"nik" validate:"required"`
	Name           string          `gorm:"not null" json:"name" validate:"required"`
	Phone          string          `gorm:"not null" json:"phone" validate:"required"`
	Gender         string          `gorm:"not null" json:"gender" validate:"required"`
	MedicalRecord  string          `gorm:"not null;unique" json:"medicalRecord"`
	CreatedAt      time.Time       `gorm:"not null" json:"createdAt"`
	UpdatedAt      time.Time       `gorm:"not null" json:"updatedAt"`
	Sessions       []Session       `json:"sessions"`
	MedicalRecords []MedicalRecord `gorm:"foreignKey:PatientMedicalRecord;references:MedicalRecord" json:"medicalRecords"`
}

type UpdatePatientPayload struct {
	ID            string `param:"patientID" validate:"required"`
	NIK           string `json:"nik" validate:"required"`
	Name          string `json:"name" validate:"required"`
	Phone         string `json:"phone" validate:"required"`
	Gender        string `json:"gender" validate:"required"`
	MedicalRecord string `json:"medicalRecord" validate:"required"`
}

type PatientIDPayload struct {
	ID string `param:"patientID" validate:"required"`
}

type PatientNIKPayload struct {
	NIK string `param:"nik" validate:"required"`
}

type AddedPatient struct {
	ID string `json:"id"`
}
