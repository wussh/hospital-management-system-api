package util

import (
	"fmt"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/postgres"
)

func InitTimes() {
	db := postgres.Connect()

	result := db.Find(&entity.Time{})

	if result.RowsAffected == 0 {
		times := []entity.Time{
			{
				Start: "08.00",
				End:   "12.00",
			},
			{
				Start: "17.00",
				End:   "20.00",
			},
			{
				Start: "08.00",
				End:   "12.00",
			},
			{
				Start: "08.00",
				End:   "12.00",
			},
			{
				Start: "08.00",
				End:   "12.00",
			},
			{
				Start: "17.00",
				End:   "20.00",
			},
			{
				Start: "08.00",
				End:   "12.00",
			},
			{
				Start: "08.00",
				End:   "12.00",
			},
			{
				Start: "08.00",
				End:   "12.00",
			},
			{
				Start: "17.00",
				End:   "20.00",
			},
			{
				Start: "08.00",
				End:   "12.00",
			},
			{
				Start: "17.00",
				End:   "20.00",
			},
		}

		for index, time := range times {
			time.ID = fmt.Sprintf("time-%d", index+1)
			time.ScheduleID = NewStringReference(fmt.Sprintf("schedule-%d", index+1))

			db.Create(&time)
		}
	}
}
