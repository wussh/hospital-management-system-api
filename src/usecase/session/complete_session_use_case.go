package session

import (
	"fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
)

type completeSessionUseCase struct {
	sessionRepository       domain.SessionRepository
	userRepository          domain.UserRepository
	jwtTokenManager         application.TokenManager
	medicalRecordRepository domain.MedicalRecordRepository
	patientRepository       domain.PatientRepository
	nanoidIDGenerator       application.IDGenerator
}

func NewCompleteSessionUseCase(
	sessionRepository domain.SessionRepository,
	userRepository domain.UserRepository,
	jwtTokenManager application.TokenManager,
	medicalRecordRepository domain.MedicalRecordRepository,
	patientRepository domain.PatientRepository,
	nanoidIDGenerator application.IDGenerator,
) domain.CompleteSessionUseCase {
	return &completeSessionUseCase{
		sessionRepository:       sessionRepository,
		userRepository:          userRepository,
		jwtTokenManager:         jwtTokenManager,
		medicalRecordRepository: medicalRecordRepository,
		patientRepository:       patientRepository,
		nanoidIDGenerator:       nanoidIDGenerator,
	}
}

func (u *completeSessionUseCase) Execute(
	sessionIDPayload entity.SessionIDPayload,
	medicalRecordPayload entity.MedicalRecord,
	authorizationHeader entity.AuthorizationHeader,
) (int, error) {
	decodedPayload, code, err := u.jwtTokenManager.DecodeAccessTokenPayload(authorizationHeader.AccessToken)
	if err != nil {
		return code, err
	}

	user, code, err := u.userRepository.GetUserByID(decodedPayload.ID)
	if err != nil {
		return code, err
	}

	session, code, err := u.sessionRepository.GetSessionByID(sessionIDPayload.ID)
	if err != nil {
		return code, err
	}

	if *session.DoctorID != user.ID {
		if user.Role != "Admin" && user.Role != "Staff" {
			return http.StatusForbidden, fmt.Errorf("restricted resource")
		}
	}

	if session.Status != "Sedang berlangsung" {
		return http.StatusBadRequest, fmt.Errorf("can't complete session which status isn't Sedang Berlangsung")
	}

	patient, code, err := u.patientRepository.GetPatientByID(*session.PatientID)
	if err != nil {
		return code, err
	}
	medicalRecordPayload.PatientMedicalRecord = &patient.MedicalRecord

	generatedID, code, err := u.nanoidIDGenerator.Generate()
	if err != nil {
		return code, err
	}
	medicalRecordPayload.ID = fmt.Sprintf("medical-record-%s", generatedID)

	if code, err := u.medicalRecordRepository.AddMedicalRecord(medicalRecordPayload); err != nil {
		return code, err
	}

	return u.sessionRepository.UpdateSessionStatusToCompleted(sessionIDPayload.ID)
}
