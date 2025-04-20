package domain

import (
	"github.com/google/uuid"
	"time"
)

type Animal struct {
	id           uuid.UUID
	enclosureId  uuid.UUID
	species      string
	name         string
	birthDate    time.Time
	gender       Gender
	favoriteFood string
	healthStatus HealthStatus
	events       []Event
}

func NewAnimal(species, name string, birthDate time.Time, gender Gender, favoriteFood string, healthStatus HealthStatus, enclosureId uuid.UUID) *Animal {
	return &Animal{
		id:           uuid.New(),
		species:      species,
		name:         name,
		birthDate:    birthDate,
		gender:       gender,
		favoriteFood: favoriteFood,
		healthStatus: healthStatus,
		enclosureId:  enclosureId,
		events:       []Event{},
	}
}

func (a *Animal) Feed(scheduledAt time.Time) time.Time {
	occurredAt := time.Now()
	a.events = append(a.events, FeedingTimeEvent{
		AnimalID:    a.id,
		FoodType:    a.favoriteFood,
		ScheduledAt: scheduledAt,
		OccurredAt:  occurredAt,
	})
	return occurredAt
}

func (a *Animal) Treat() {
	a.healthStatus = Healthy
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

func (a *Animal) ID() uuid.UUID              { return a.id }
func (a *Animal) EnclosureID() uuid.UUID     { return a.enclosureId }
func (a *Animal) Species() string            { return a.species }
func (a *Animal) Name() string               { return a.name }
func (a *Animal) BirthDate() time.Time       { return a.birthDate }
func (a *Animal) Gender() Gender             { return a.gender }
func (a *Animal) FavoriteFood() string       { return a.favoriteFood }
func (a *Animal) HealthStatus() HealthStatus { return a.healthStatus }

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
