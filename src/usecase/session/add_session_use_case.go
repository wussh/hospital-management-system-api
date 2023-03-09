package session

import (
	"fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
)

type addSessionUseCase struct {
	sessionRepository  domain.SessionRepository
	patientRepository  domain.PatientRepository
	clinicRepository   domain.ClinicRepository
	userRepository     domain.UserRepository
	scheduleRepository domain.ScheduleRepository
	nanoidIDGenerator  application.IDGenerator
}

func NewAddSessionUseCase(
	sessionRepository domain.SessionRepository,
	patientRepository domain.PatientRepository,
	clinicRepository domain.ClinicRepository,
	userRepository domain.UserRepository,
	scheduleRepository domain.ScheduleRepository,
	nanoidIDGenerator application.IDGenerator,
) domain.AddSessionUseCase {
	return &addSessionUseCase{
		sessionRepository:  sessionRepository,
		patientRepository:  patientRepository,
		clinicRepository:   clinicRepository,
		userRepository:     userRepository,
		scheduleRepository: scheduleRepository,
		nanoidIDGenerator:  nanoidIDGenerator,
	}
}

func (u *addSessionUseCase) Execute(payload entity.Session) (entity.Session, int, error) {
	if _, code, err := u.patientRepository.GetPatientByID(*payload.PatientID); err != nil {
		return entity.Session{}, code, err
	}

	if _, code, err := u.clinicRepository.GetClinicByID(*payload.ClinicID); err != nil {
		return entity.Session{}, code, err
	}

	if _, code, err := u.userRepository.GetUserDoctorByID(*payload.DoctorID); err != nil {
		return entity.Session{}, code, err
	}

	if _, code, err := u.scheduleRepository.GetScheduleByID(*payload.ScheduleID); err != nil {
		return entity.Session{}, code, err
	}

	queue, code, err := u.sessionRepository.GetSessionLastQueue(*payload.ScheduleID, payload.Date)
	if err != nil {
		return entity.Session{}, code, err
	}

	payload.Queue = queue + 1
	payload.QueueCode = fmt.Sprintf("%s-%d", *payload.ClinicID, payload.Queue)

	generatedID, code, err := u.nanoidIDGenerator.Generate()
	if err != nil {
		return entity.Session{}, code, err
	}
	payload.ID = fmt.Sprintf("session-%s", generatedID)

	payload.Status = "Dalam antrian"

	session, code, err := u.sessionRepository.AddSession(payload)
	if err != nil {
		return entity.Session{}, code, err
	}

	return session, http.StatusOK, nil
}
