package domain

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type FeedingSchedule struct {
	id          uuid.UUID
	AnimalId    uuid.UUID
	FoodType    FoodType
	ScheduledAt time.Time
	Occurred    bool
}

func NewFeedingSchedule(animalId uuid.UUID, foodType FoodType, scheduledAt time.Time) FeedingSchedule {
	return FeedingSchedule{
		id:          uuid.New(),
		AnimalId:    animalId,
		FoodType:    foodType,
		ScheduledAt: scheduledAt,
		Occurred:    false,
	}
}

func (fs *FeedingSchedule) MarkAsOccurred(occurredAt time.Time) error {
	if fs.Occurred {
		return fmt.Errorf("feeding schedule already marked as occurred")
	}
	fs.Occurred = true
	return nil
}

func (fs *FeedingSchedule) ID() uuid.UUID {
	return fs.id
}
