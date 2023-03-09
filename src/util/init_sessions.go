package util

import (
	"fmt"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/postgres"
)

func InitSessions() {
	db := postgres.Connect()

	result := db.Find(&entity.Session{})

	if result.RowsAffected == 0 {
		sessions := []entity.Session{
			{
				Status: "Selesai",
			},
			{
				Status: "Dibatalkan",
			},
		}

		for index, session := range sessions {
			session.ID = fmt.Sprintf("session-%d", index+1)
			session.PatientID = NewStringReference(fmt.Sprintf("patient-%d", index+1))
			session.ClinicID = NewStringReference(fmt.Sprintf("clinic-%d", index+1))
			session.DoctorID = NewStringReference(fmt.Sprintf("doctor-%d", index+1))
			session.ScheduleID = NewStringReference(fmt.Sprintf("schedule-%d", index+1))
			session.Complaint = "sebuah keluhan"
			session.Queue = 1
			session.QueueCode = fmt.Sprintf("%s-%d", *session.ClinicID, session.Queue)
			session.Date = "2022/7/6"

			db.Create(&session)

			if index == 0 {
				record := entity.MedicalRecord{
					ID:                   fmt.Sprintf("medical-record-%d", index+1),
					SessionID:            &session.ID,
					PatientMedicalRecord: NewStringReference(fmt.Sprintf("record-%d", index+1)),
					Type:                 "Rawat Jalan",
					History:              "Tidak ada",
					Diagnosis:            "Sembuh",
					Height:               "165",
					Weight:               "60",
					Systole:              "70",
					Diastole:             "120",
					Temperature:          "35",
					Status:               "Pulang",
				}

				db.Create(&record)
			}
		}
	}
}
