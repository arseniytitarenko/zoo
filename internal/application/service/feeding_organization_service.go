package service

import (
	"github.com/google/uuid"
	"time"
	"zoo/internal/application/dto"
	"zoo/internal/application/errs"
	"zoo/internal/application/port/out"
	"zoo/internal/domain"
)

type FeedingOrganizationService struct {
	animalRepo          out.AnimalRepository
	feedingScheduleRepo out.FeedingScheduleRepository
	eventDispatcher     out.IEventDispatcher
}

func NewFeedingOrganizationService(animalRepo out.AnimalRepository, feedingScheduleRepo out.FeedingScheduleRepository, dispatcher out.IEventDispatcher) *FeedingOrganizationService {
	return &FeedingOrganizationService{animalRepo: animalRepo, feedingScheduleRepo: feedingScheduleRepo, eventDispatcher: dispatcher}
}

func (s *FeedingOrganizationService) GetAllFeedingSchedules() []dto.FeedingScheduleResponse {
	schedules := s.feedingScheduleRepo.GetAll()
	return newFeedingScheduleResponses(schedules)
}

func (s *FeedingOrganizationService) GetAnimalFeedingSchedulesByID(animalId string) ([]dto.FeedingScheduleResponse, error) {
	animalUUID, err := uuid.Parse(animalId)
	if err != nil {
		return nil, errs.ErrInvalidID
	}
	schedules := s.feedingScheduleRepo.GetAllByAnimalID(animalUUID)
	return newFeedingScheduleResponses(schedules), nil
}

func (s *FeedingOrganizationService) NewFeedingSchedule(req *dto.NewFeedingScheduleRequest) (*dto.FeedingScheduleResponse, error) {
	scheduledAt, err := time.Parse("02.01.2006 15:04", req.ScheduledAt)
	if err != nil {
		return nil, errs.ErrInvalidTime
	}
	animalUUID, err := uuid.Parse(req.AnimalID)
	if err != nil {
		return nil, errs.ErrInvalidID
	}
	animal, ok := s.animalRepo.GetByID(animalUUID)
	if !ok {
		return nil, errs.ErrAnimalNotFound
	}
	fs := domain.NewFeedingSchedule(animal.ID(), req.FoodType, scheduledAt)
	s.feedingScheduleRepo.Save(fs)
	return newFeedingScheduleResponse(fs), nil
}

func (s *FeedingOrganizationService) FeedByScheduleId(scheduleId string) error {
	scheduleUUID, err := uuid.Parse(scheduleId)
	if err != nil {
		return errs.ErrInvalidID
	}
	schedule, ok := s.feedingScheduleRepo.GetByID(scheduleUUID)
	if !ok {
		return errs.ErrScheduleNotFound
	}
	animal, ok := s.animalRepo.GetByID(schedule.AnimalID())
	if !ok {
		return errs.ErrAnimalNotFound
	}
	if schedule.IsOccurred() {
		return errs.ErrFeedingAlreadyOccurred
	}

	occurredAt := animal.Feed(schedule.ScheduledAt())
	err = schedule.MarkAsOccurred(occurredAt)
	if err != nil {
		return err
	}

	events := animal.PullEvents()
	for _, event := range events {
		s.eventDispatcher.Dispatch(event)
	}
	return nil
}

func newFeedingScheduleResponses(schedules []domain.FeedingSchedule) []dto.FeedingScheduleResponse {
	result := make([]dto.FeedingScheduleResponse, len(schedules))
	for i, schedule := range schedules {
		result[i] = *newFeedingScheduleResponse(&schedule)
	}
	return result
}

func newFeedingScheduleResponse(schedule *domain.FeedingSchedule) *dto.FeedingScheduleResponse {
	return &dto.FeedingScheduleResponse{
		ID:          schedule.ID(),
		AnimalID:    schedule.AnimalID(),
		FoodType:    schedule.FoodType(),
		ScheduledAt: schedule.ScheduledAt().Format("02.01.2006 15:04"),
		OccurredAt:  schedule.OccurredAt().Format("02.01.2006 15:04"),
	}
}
