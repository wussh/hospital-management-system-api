package entity

import (
	"mime/multipart"
	"time"
)

type User struct {
	ID         string     `gorm:"not null,primaryKey,index:idx_id" json:"id"`
	Name       string     `gorm:"not null" json:"name" validate:"required"`
	Speciality string     `json:"speciality"`
	Phone      string     `gorm:"not null" json:"phone" validate:"required"`
	Password   string     `gorm:"not null" json:"-" validate:"required"`
	Role       string     `gorm:"not null,index:idx_role" json:"role" validate:"required"`
	License    string     `json:"license"`
	Address    string     `gorm:"not null" json:"address" validate:"required"`
	BirthPlace string     `gorm:"not null" json:"birthPlace" validate:"required"`
	BirthDate  string     `gorm:"not null" json:"birthDate" validate:"required"`
	Religion   string     `gorm:"not null" json:"religion" validate:"required"`
	Avatar     string     `json:"avatar"`
	ClinicID   *string    `json:"clinicID"`
	CreatedAt  time.Time  `gorm:"not null" json:"createdAt"`
	UpdatedAt  time.Time  `gorm:"not null" json:"updatedAt"`
	Clinic     Clinic     `json:"clinic"`
	Schedules  []Schedule `json:"schedules"`
}

type AuthenticationPayload struct {
	ID string
}

type LoginPayload struct {
	ID       string `json:"id" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AddedUser struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateAvatarPayload struct {
	ID     string `param:"userID" validate:"required"`
	Avatar *multipart.FileHeader
}

type UpdateAvatarLocationPayload struct {
	ID     string
	Avatar string
}

type UpdatedAvatar struct {
	Avatar string `json:"avatar"`
}

type UserIDPayload struct {
	ID string `param:"userID" validate:"required"`
}
