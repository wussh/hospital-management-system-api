package util

import (
	"fmt"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/postgres"
)

func InitSchedules() {
	db := postgres.Connect()

	result := db.Find(&entity.Schedule{})

	if result.RowsAffected == 0 {
		schedules := []entity.Schedule{
			{
				DayID:  NewStringReference("day-2"),
				UserID: NewStringReference("doctor-1"),
			},
			{
				DayID:  NewStringReference("day-2"),
				UserID: NewStringReference("doctor-1"),
			},
			{
				DayID:  NewStringReference("day-3"),
				UserID: NewStringReference("doctor-2"),
			},
			{
				DayID:  NewStringReference("day-1"),
				UserID: NewStringReference("doctor-2"),
			},
			{
				DayID:  NewStringReference("day-4"),
				UserID: NewStringReference("doctor-3"),
			},
			{
				DayID:  NewStringReference("day-4"),
				UserID: NewStringReference("doctor-3"),
			},
			{
				DayID:  NewStringReference("day-5"),
				UserID: NewStringReference("doctor-4"),
			},
			{
				DayID:  NewStringReference("day-1"),
				UserID: NewStringReference("doctor-4"),
			},
			{
				DayID:  NewStringReference("day-6"),
				UserID: NewStringReference("doctor-5"),
			},
			{
				DayID:  NewStringReference("day-6"),
				UserID: NewStringReference("doctor-5"),
			},
			{
				DayID:  NewStringReference("day-1"),
				UserID: NewStringReference("doctor-6"),
			},
			{
				DayID:  NewStringReference("day-1"),
				UserID: NewStringReference("doctor-6"),
			},
		}

		for index, schedule := range schedules {
			schedule.ID = fmt.Sprintf("schedule-%d", index+1)

			db.Create(&schedule)
		}
	}
}
