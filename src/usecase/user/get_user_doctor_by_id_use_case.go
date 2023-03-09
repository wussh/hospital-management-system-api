package user

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type getUserDoctorByIDUseCase struct {
	userRepository     domain.UserRepository
	scheduleRepository domain.ScheduleRepository
	dayRepository      domain.DayRepository
	timeRepository     domain.TimeRepository
	clinicRepository   domain.ClinicRepository
}

func NewGetUserDoctorByIDUseCase(
	userRepository domain.UserRepository,
	scheduleRepository domain.ScheduleRepository,
	dayRepository domain.DayRepository,
	timeRepository domain.TimeRepository,
	clinicRepository domain.ClinicRepository,
) domain.GetUserDoctorByIDUseCase {
	return &getUserDoctorByIDUseCase{
		userRepository:     userRepository,
		scheduleRepository: scheduleRepository,
		dayRepository:      dayRepository,
		timeRepository:     timeRepository,
		clinicRepository:   clinicRepository,
	}
}

func (u *getUserDoctorByIDUseCase) Execute(payload entity.UserIDPayload) (entity.User, int, error) {
	user, code, err := u.userRepository.GetUserDoctorByID(payload.ID)
	if err != nil {
		return entity.User{}, code, err
	}

	schedules, code, err := u.scheduleRepository.GetSchedulesByDoctorID(payload.ID)
	if err != nil {
		return entity.User{}, code, err
	}

	for i := 0; i < len(schedules); i++ {
		day, code, err := u.dayRepository.GetDayByID(*schedules[i].DayID)
		if err != nil {
			return entity.User{}, code, err
		}

		time, code, err := u.timeRepository.GetTimeByScheduleID(schedules[i].ID)
		if err != nil {
			return entity.User{}, code, err
		}

		schedules[i].Day = day
		schedules[i].Time = time
	}
	user.Schedules = schedules

	clinic, code, err := u.clinicRepository.GetClinicByID(*user.ClinicID)
	if err != nil {
		return entity.User{}, code, err
	}
	user.Clinic = clinic

	return user, code, err
}
