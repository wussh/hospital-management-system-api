package util

import (
	"fmt"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/postgres"
	"golang.org/x/crypto/bcrypt"
)

func InitDoctors() {
	db := postgres.Connect()

	result := db.Where("role = ?", "Doctor").Find(&entity.User{})

	if result.RowsAffected == 0 {
		doctors := []entity.User{
			{
				Name:       "dr. Jane",
				Password:   "drjane",
				Speciality: "Umum",
			},
			{
				Name:       "dr. Richard",
				Password:   "drrichard",
				Speciality: "Mata",
			},
			{
				Name:       "dr. Peter",
				Password:   "drpeter",
				Speciality: "Jantung",
			},
			{
				Name:       "dr. Wanda",
				Password:   "drwanda",
				Speciality: "Anak",
			},
			{
				Name:       "dr. Strange",
				Password:   "drstrange",
				Speciality: "Telinga",
			},
			{
				Name:       "dr. Kamala",
				Password:   "drkamala",
				Speciality: "Gigi",
			},
		}

		for index, doctor := range doctors {
			doctor.ID = fmt.Sprintf("doctor-%d", index+1)

			hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(doctor.Password), bcrypt.DefaultCost)
			doctor.Password = string(hashedPassword)

			doctor.Phone = "08123456789"
			doctor.Role = "Doctor"
			doctor.License = fmt.Sprintf("doctor-license-%d", index+1)
			doctor.Address = "Kebon Jeruk, Jakarta Barat"
			doctor.BirthPlace = "Jakarta"
			doctor.BirthDate = "1-1-1970"
			doctor.Religion = "Islam"

			doctor.ClinicID = NewStringReference(fmt.Sprintf("clinic-%d", index+1))

			db.Create(&doctor)
		}
	}
}
