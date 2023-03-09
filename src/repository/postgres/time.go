package postgres

import (
	"fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"gorm.io/gorm"
)

type timeRepository struct {
	db *gorm.DB
}

func NewTimeRepository(db *gorm.DB) domain.TimeRepository {
	return &timeRepository{
		db: db,
	}
}

func (r *timeRepository) GetTimeByScheduleID(scheduleID string) (entity.Time, int, error) {
	time := entity.Time{}
	result := r.db.Where("schedule_id = ?", scheduleID).First(&time)

	if result.RowsAffected == 0 {
		return entity.Time{}, http.StatusNotFound, fmt.Errorf("time not found")
	}

	return time, http.StatusOK, nil
}
