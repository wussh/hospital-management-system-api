package util

import (
	"fmt"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/postgres"
)

func InitPatients() {
	db := postgres.Connect()

	result := db.Find(&entity.Patient{})

	if result.RowsAffected == 0 {
		patients := []entity.Patient{
			{
				NIK:  "4517092810000001",
				Name: "John",
			},
			{
				NIK:  "4517092810000002",
				Name: "Doe",
			},
		}

		for index, patient := range patients {
			patient.ID = fmt.Sprintf("patient-%d", index+1)

			patient.Phone = "08123456789"
			patient.Gender = "Laki-Laki"
			patient.MedicalRecord = fmt.Sprintf("record-%d", index+1)

			db.Create(&patient)
		}
	}
}
