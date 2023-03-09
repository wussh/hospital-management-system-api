package util

import (
	"fmt"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/postgres"
)

func InitDays() {
	db := postgres.Connect()

	result := db.Find(&entity.Day{})

	if result.RowsAffected == 0 {
		days := []entity.Day{
			{Name: "Senin"},
			{Name: "Selasa"},
			{Name: "Rabu"},
			{Name: "Kamis"},
			{Name: "Jum'at"},
			{Name: "Sabtu"},
			{Name: "Minggu"},
		}

		for index, day := range days {
			day.ID = fmt.Sprintf("day-%d", index+1)
			day.Order = index

			db.Create(&day)
		}
	}
}
