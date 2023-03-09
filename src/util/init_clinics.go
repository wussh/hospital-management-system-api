package util

import (
	"fmt"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/postgres"
)

func InitClinics() {
	db := postgres.Connect()

	result := db.Find(&entity.Clinic{})

	if result.RowsAffected == 0 {
		clinics := []entity.Clinic{
			{Name: "Umum"},
			{Name: "Mata"},
			{Name: "Jantung"},
			{Name: "Anak"},
			{Name: "THT"},
			{Name: "Gigi"},
		}

		for index, clinic := range clinics {
			clinic.ID = fmt.Sprintf("clinic-%d", index+1)

			db.Create(&clinic)
		}
	}
}
