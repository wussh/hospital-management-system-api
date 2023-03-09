package entity

import "time"

type Day struct {
	ID        string    `gorm:"not null" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Order     int       `gorm:"not null" json:"order"`
	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
}
