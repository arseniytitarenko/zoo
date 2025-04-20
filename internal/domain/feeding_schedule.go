package domain

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type FeedingSchedule struct {
	id          uuid.UUID
	animalId    uuid.UUID
	foodType    string
	scheduledAt time.Time
	occurredAt  *time.Time
}

func NewFeedingSchedule(animalId uuid.UUID, foodType string, scheduledAt time.Time) *FeedingSchedule {
	return &FeedingSchedule{
		id:          uuid.New(),
		animalId:    animalId,
		foodType:    foodType,
		scheduledAt: scheduledAt,
	}
}

func (fs *FeedingSchedule) MarkAsOccurred(occurredAt time.Time) error {
	if fs.IsOccurred() {
		return fmt.Errorf("feeding schedule already marked as occurred")
	}
	fs.occurredAt = &occurredAt
	return nil
}

func (fs *FeedingSchedule) IsOccurred() bool {
	return fs.occurredAt != nil
}

func (fs *FeedingSchedule) ChangeScheduleTime(newTime time.Time) {
	fs.scheduledAt = newTime
}

func (fs *FeedingSchedule) ID() uuid.UUID {
	return fs.id
}

func (fs *FeedingSchedule) AnimalID() uuid.UUID {
	return fs.animalId
}

func (fs *FeedingSchedule) FoodType() string {
	return fs.foodType
}

func (fs *FeedingSchedule) ScheduledAt() time.Time {
	return fs.scheduledAt
}

func (fs *FeedingSchedule) OccurredAt() *time.Time {
	return fs.occurredAt
}
