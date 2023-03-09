package entity

import (
	"time"
)

type Clinic struct {
	ID        string    `gorm:"not null,primaryKey,index:idx_id" json:"id"`
	Name      string    `gorm:"not null" json:"name" validate:"required"`
	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
	Doctors   []User    `json:"doctors"`
}

type UpdateClinicPayload struct {
	ID   string `param:"clinicID" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type ClinicIDPayload struct {
	ID string `param:"clinicID" validate:"required"`
}
