package entity

import "time"

type Schedule struct {
	ID        string    `gorm:"not null" json:"id"`
	DayID     *string   `gorm:"not null" json:"dayID" validate:"required"`
	UserID    *string   `gorm:"not null" json:"userID" validate:"required"`
	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
	Day       Day       `json:"day"`
	Time      Time      `json:"time"`
}

type UpdateSchedulePayload struct {
	ID     string `param:"scheduleID" validate:"required"`
	DayID  string `json:"dayID" validate:"required"`
	UserID string `json:"userID" validate:"required"`
}

type ScheduleIDPayload struct {
	ID     string `param:"scheduleID" validate:"required"`
	UserID string `param:"userID" validate:"required"`
}
