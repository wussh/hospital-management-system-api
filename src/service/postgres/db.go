package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_TIMEZONE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to postgres database")
	}

	return db
}

func InitMigration() {
	db := Connect()

	db.AutoMigrate(
		&entity.Authentication{},
		&entity.Clinic{},
		&entity.User{},
		&entity.Day{},
		&entity.Schedule{},
		&entity.Time{},
		&entity.Patient{},
		&entity.Session{},
		&entity.MedicalRecord{},
		// &entity.Staff{},
		// &entity.Doctor{},
	)
}
