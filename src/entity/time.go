package entity

import "time"

type Time struct {
	ID         string    `gorm:"not null" json:"ID"`
	Start      string    `gorm:"not null" json:"start" validation:"required"`
	End        string    `gorm:"not null" json:"end" validation:"required"`
	ScheduleID *string   `gorm:"not null" json:"scheduleID" validation:"required"`
	CreatedAt  time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"not null" json:"updatedAt"`
}
