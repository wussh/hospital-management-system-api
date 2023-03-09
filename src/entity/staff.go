package entity

import "gorm.io/gorm"

type Staff struct {
	gorm.Model
	StaffType string `gorm:"not null" json:"staffType" validate:"required"`
	Name      string `gorm:"not null" json:"name" validate:"required"`
	Phone     string `gorm:"not null" json:"phone" validate:"required"`
	Email     string `gorm:"not null;unique" json:"email" validate:"required,email"`
	Password  string `gorm:"not null" json:"password" validate:"required"`
	Token     string `gorm:"not null" json:"token" validate:"required"`
}
