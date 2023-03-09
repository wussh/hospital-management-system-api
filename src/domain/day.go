package domain

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type DayRepository interface {
	GetDayByID(id string) (entity.Day, int, error)
}
