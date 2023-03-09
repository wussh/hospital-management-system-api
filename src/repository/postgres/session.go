package postgres

import (
	"fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/util"
	"gorm.io/gorm"
)

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) domain.SessionRepository {
	return &sessionRepository{
		db: db,
	}
}

func (r *sessionRepository) AddSession(payload entity.Session) (entity.Session, int, error) {
	result := r.db.Create(&payload)

	if result.RowsAffected == 0 {
		return entity.Session{}, http.StatusInternalServerError, result.Error
	}

	return payload, http.StatusOK, nil
}

func (r *sessionRepository) GetSessionLastQueue(scheduleID string, date string) (int, int, error) {
	lastQueue := util.NewIntReference(0)

	result := r.db.Find(&entity.Session{})
	if result.RowsAffected != 0 {
		row := r.db.Table("sessions").Where("schedule_id = ? AND date = ?", scheduleID, date).Select("MAX(queue)").Row()
		err := row.Scan(&lastQueue)

		if err != nil {
			return -1, http.StatusInternalServerError, err
		}
	}

	if result.Error != nil {
		return -1, http.StatusInternalServerError, result.Error
	}

	if lastQueue == nil {
		lastQueue = util.NewIntReference(0)
	}

	return *lastQueue, http.StatusOK, nil
}

func (r *sessionRepository) GetSessionByID(id string) (entity.Session, int, error) {
	session := entity.Session{}
	result := r.db.Where("id = ?", id).First(&session)

	if result.RowsAffected == 0 {
		return entity.Session{}, http.StatusNotFound, fmt.Errorf("session not found")
	}

	return session, http.StatusOK, nil
}

func (r *sessionRepository) GetSessionByPatientID(patientID string) ([]entity.Session, int, error) {
	sessions := []entity.Session{}
	r.db.Where("patient_id = ?", patientID).Find(&sessions)

	return sessions, http.StatusOK, nil
}

func (r *sessionRepository) GetSessions() ([]entity.Session, int, error) {
	sessions := []entity.Session{}
	r.db.Find(&sessions)

	return sessions, http.StatusOK, nil
}

func (r *sessionRepository) GetSessionsByDoctorID(doctorID string) ([]entity.Session, int, error) {
	sessions := []entity.Session{}
	r.db.Where("doctor_id = ?", doctorID).Find(&sessions)

	return sessions, http.StatusOK, nil
}

func (r *sessionRepository) GetQueuedSessionsByDoctorID(doctorID string) ([]entity.Session, int, error) {
	sessions := []entity.Session{}
	r.db.Where("doctor_id = ? AND status = ?", doctorID, "Dalam antrian").Find(&sessions)

	return sessions, http.StatusOK, nil
}

func (r *sessionRepository) GetCompletedSessionsByDoctorID(doctorID string) ([]entity.Session, int, error) {
	sessions := []entity.Session{}
	r.db.Where("doctor_id = ? AND status = ?", doctorID, "Selesai").Find(&sessions)

	return sessions, http.StatusOK, nil
}

func (r *sessionRepository) GetCancelledSessionsByDoctorID(doctorID string) ([]entity.Session, int, error) {
	sessions := []entity.Session{}
	r.db.Where("doctor_id = ? AND status = ?", doctorID, "Dibatalkan").Find(&sessions)

	return sessions, http.StatusOK, nil
}

func (r *sessionRepository) UpdateSessionStatusToCompleted(id string) (int, error) {
	session, code, err := r.GetSessionByID(id)
	if err != nil {
		return code, err
	}
	session.Status = "Selesai"

	r.db.Save(&session)

	return http.StatusOK, nil
}

func (r *sessionRepository) UpdateSessionStatusToCancelled(id string) (int, error) {
	session, code, err := r.GetSessionByID(id)
	if err != nil {
		return code, err
	}
	session.Status = "Dibatalkan"

	r.db.Save(&session)

	return http.StatusOK, nil
}

func (r *sessionRepository) UpdateSessionStatusToActive(id string) (int, error) {
	session, code, err := r.GetSessionByID(id)
	if err != nil {
		return code, err
	}
	session.Status = "Sedang berlangsung"

	r.db.Save(&session)

	return http.StatusOK, nil
}

func (r *sessionRepository) VerifyNoActiveSession() (int, error) {
	result := r.db.Where("status = ?", "Sedang berlangsung").Find(&entity.Session{})

	if result.RowsAffected > 0 {
		return http.StatusBadRequest, fmt.Errorf("there is an active session. complete it first")
	}

	return http.StatusOK, nil
}
