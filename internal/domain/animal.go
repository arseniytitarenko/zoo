package domain

import (
	"github.com/google/uuid"
	"time"
)

type Animal struct {
	Id           uuid.UUID
	EnclosureId  uuid.UUID
	Species      string
	Name         string
	BirthDate    time.Time
	Gender       Gender
	FavoriteFood string
	HealthStatus HealthStatus
	events       []Event
}

func (a *Animal) Feed(scheduledAt time.Time) {
	a.events = append(a.events, FeedingTimeEvent{
		AnimalID:    a.Id,
		FoodType:    a.FavoriteFood,
		ScheduledAt: scheduledAt,
		OccurredAt:  time.Now(),
	})
}

func (a *Animal) Treat() {
	a.HealthStatus = Healthy
}

func (a *Animal) MoveTo(newEnclosureId uuid.UUID) {
	old := a.EnclosureId
	a.EnclosureId = newEnclosureId

	a.events = append(a.events, AnimalMovedEvent{
		AnimalID:        a.Id,
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
