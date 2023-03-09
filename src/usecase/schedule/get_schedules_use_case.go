package schedule

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type getSchedulesUseCase struct {
	scheduleRepository domain.ScheduleRepository
}

func NewGetSchedulesUseCase(scheduleRepository domain.ScheduleRepository) domain.GetSchedulesUseCase {
	return &getSchedulesUseCase{
		scheduleRepository: scheduleRepository,
	}
}

func (u *getSchedulesUseCase) Execute() ([]entity.Schedule, int, error) {
	return u.scheduleRepository.GetSchedules()
}
