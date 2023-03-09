package util

import (
	"os"
	"time"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/postgres"

	"golang.org/x/crypto/bcrypt"
)

func CreateAdminUser() {
	db := postgres.Connect()

	result := db.Where("role = ?", "Admin").First(&entity.User{})

	if result.RowsAffected == 0 {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(os.Getenv("ADMIN_PASSWORD")), bcrypt.DefaultCost)
		createdAt := time.Now()

		db.Create(&entity.User{
			ID:         os.Getenv("ADMIN_ID"),
			Name:       "Admin Simars",
			Speciality: "Admin",
			Phone:      os.Getenv("ADMIN_PHONE"),
			Password:   string(hashedPassword),
			Role:       "Admin",
			Address:    "Simars",
			BirthPlace: "Jakarta",
			BirthDate:  "1-1-1970",
			Religion:   "Admin",
			CreatedAt:  createdAt,
			UpdatedAt:  createdAt,
		})
	}
}
