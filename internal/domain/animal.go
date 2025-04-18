package domain

import (
	"github.com/google/uuid"
	"time"
)

type Animal struct {
	id           uuid.UUID
	enclosureId  uuid.UUID
	Species      string
	Name         string
	DateOfBirth  time.Time
	Gender       Gender
	FavoriteFood FoodType
	Status       HealthStatus
	events       []Event
}

func (a *Animal) Feed(scheduledAt time.Time) {
	a.events = append(a.events, FeedingTimeEvent{
		AnimalID:    a.id,
		FoodType:    a.FavoriteFood,
		ScheduledAt: scheduledAt,
		OccurredAt:  time.Now(),
	})
}

func (a *Animal) Treat() {
	a.Status = Healthy
}

func (a *Animal) MoveTo(newEnclosureId uuid.UUID) {
	old := a.enclosureId
	a.enclosureId = newEnclosureId

	a.events = append(a.events, AnimalMovedEvent{
		AnimalID:        a.id,
		FromEnclosureID: old,
		ToEnclosureID:   newEnclosureId,
		OccurredAt:      time.Now(),
	})
}

func (a *Animal) PullEvents() []Event {
	evs := a.events
	a.events = nil
	return evs
}

func (a *Animal) ID() uuid.UUID {
	return a.id
}

func (a *Animal) EnclosureID() uuid.UUID {
	return a.enclosureId
}

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

type HealthStatus string

const (
	Healthy HealthStatus = "Healthy"
	Sick    HealthStatus = "Sick"
)

type FoodType string
