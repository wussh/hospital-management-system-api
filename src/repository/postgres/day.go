package postgres

import (
	"fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"gorm.io/gorm"
)

type dayRepository struct {
	db *gorm.DB
}

func NewDayRepository(db *gorm.DB) domain.DayRepository {
	return &dayRepository{
		db: db,
	}
}

func (r *dayRepository) GetDayByID(id string) (entity.Day, int, error) {
	day := entity.Day{}
	result := r.db.Where("id = ?", id).First(&day)

	if result.RowsAffected == 0 {
		return entity.Day{}, http.StatusNotFound, fmt.Errorf("day not found")
	}

	return day, http.StatusOK, nil
}
