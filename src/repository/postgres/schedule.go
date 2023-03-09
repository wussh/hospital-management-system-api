package postgres

import (
	"fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"gorm.io/gorm"
)

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) domain.ScheduleRepository {
	return &scheduleRepository{
		db: db,
	}
}

func (r *scheduleRepository) AddSchedule(payload entity.Schedule) (int, error) {
	result := r.db.Create(&payload)

	if result.RowsAffected == 0 {
		return http.StatusInternalServerError, result.Error
	}

	return http.StatusOK, nil
}

func (r *scheduleRepository) GetSchedules() ([]entity.Schedule, int, error) {
	schedules := []entity.Schedule{}
	r.db.Find(&schedules)

	return schedules, http.StatusOK, nil
}

func (r *scheduleRepository) GetScheduleByID(id string) (entity.Schedule, int, error) {
	schedule := entity.Schedule{}
	result := r.db.Where("id = ?", id).First(&schedule)

	if result.RowsAffected == 0 {
		return entity.Schedule{}, http.StatusNotFound, fmt.Errorf("schedule not found")
	}

	return schedule, http.StatusOK, nil
}

func (r *scheduleRepository) UpdateScheduleByID(payload entity.UpdateSchedulePayload) (int, error) {
	schedule, code, err := r.GetScheduleByID(payload.ID)
	if err != nil {
		return code, err
	}

	*schedule.DayID = payload.DayID
	*schedule.UserID = payload.UserID

	result := r.db.Save(&schedule)

	if result.Error != nil {
		return http.StatusInternalServerError, result.Error
	}

	return http.StatusOK, nil
}

func (r *scheduleRepository) DeleteScheduleByID(id string) (int, error) {
	result := r.db.Where("id = ?", id).Delete(&entity.Schedule{})

	if result.RowsAffected == 0 {
		return http.StatusNotFound, fmt.Errorf("schedule not found")
	}

	return http.StatusOK, nil
}

func (r *scheduleRepository) GetSchedulesByDoctorID(doctorID string) ([]entity.Schedule, int, error) {
	schedules := []entity.Schedule{}
	r.db.Where("user_id = ?", doctorID).Find(&schedules)

	return schedules, http.StatusOK, nil
}
