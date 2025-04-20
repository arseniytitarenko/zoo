package in

import "zoo/internal/application/dto"

type FeedingOrganizationUseCase interface {
	GetAllFeedingSchedules() []dto.FeedingScheduleResponse
	GetAnimalFeedingSchedulesByID(animalId string) ([]dto.FeedingScheduleResponse, error)
	NewFeedingSchedule(req *dto.NewFeedingScheduleRequest) (*dto.FeedingScheduleResponse, error)
	FeedByScheduleId(scheduleId string) (*dto.FeedingScheduleResponse, error)
}
