package domain

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type TimeRepository interface {
	GetTimeByScheduleID(scheduleID string) (entity.Time, int, error)
}
