package domain

import (
	"github.com/google/uuid"
	"time"
)

type Event interface {
	EventName() EventName
	OccurredAtTime() time.Time
}

type AnimalMovedEvent struct {
	AnimalID        uuid.UUID
	FromEnclosureID uuid.UUID
	ToEnclosureID   uuid.UUID
	OccurredAt      time.Time
}

func (e AnimalMovedEvent) EventName() EventName {
	return AnimalMoved
}

func (e AnimalMovedEvent) OccurredAtTime() time.Time {
	return e.OccurredAt
}

type FeedingTimeEvent struct {
	AnimalID    uuid.UUID
	FoodType    FoodType
	ScheduledAt time.Time
	OccurredAt  time.Time
}

func (e FeedingTimeEvent) EventName() EventName {
	return FeedingTime
}

func (e FeedingTimeEvent) OccurredAtTime() time.Time {
	return e.OccurredAt
}

type EventName string

const (
	FeedingTime EventName = "FeedingTime"
	AnimalMoved EventName = "AnimalMoved"
)
