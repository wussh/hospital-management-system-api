package entity

import "time"

type Session struct {
	ID         string    `gorm:"not null" json:"id"`
	PatientID  *string   `gorm:"not null" json:"patientID" validate:"required"`
	ClinicID   *string   `gorm:"not null" json:"clinicID" validate:"required"`
	DoctorID   *string   `gorm:"not null" json:"doctorID" validate:"required"`
	ScheduleID *string   `gorm:"not null" json:"scheduleID" validate:"required"`
	Complaint  string    `gorm:"not null" json:"complaint" validate:"required"`
	Queue      int       `gorm:"not null" json:"queue"`
	QueueCode  string    `gorm:"not null" json:"queueCode"`
	Status     string    `gorm:"not null" json:"status"`
	Date       string    `gorm:"not null" json:"date" validate:"required"`
	CreatedAt  time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"not null" json:"updatedAt"`
}

type GetSessionParams struct {
	Status string `query:"status"`
}

type SessionIDPayload struct {
	ID string `param:"sessionID"`
}
