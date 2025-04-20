package repository

import (
	"github.com/google/uuid"
	"sync"
	"zoo/internal/domain"
)

type InMemoryFeedingScheduleRepo struct {
	data map[uuid.UUID]*domain.FeedingSchedule
	mu   sync.RWMutex
}

func NewInMemoryFeedingScheduleRepo() *InMemoryFeedingScheduleRepo {
	return &InMemoryFeedingScheduleRepo{
		data: make(map[uuid.UUID]*domain.FeedingSchedule),
	}
}

func (r *InMemoryFeedingScheduleRepo) Save(feedingSchedule *domain.FeedingSchedule) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[feedingSchedule.ID()] = feedingSchedule
}

func (r *InMemoryFeedingScheduleRepo) Delete(id uuid.UUID) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.data, id)
}

func (r *InMemoryFeedingScheduleRepo) GetAll() []domain.FeedingSchedule {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var feedingSchedules []domain.FeedingSchedule
	for _, a := range r.data {
		feedingSchedules = append(feedingSchedules, *a)
	}
	return feedingSchedules
}

func (r *InMemoryFeedingScheduleRepo) GetByID(id uuid.UUID) (*domain.FeedingSchedule, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	feedingSchedule, exists := r.data[id]
	return feedingSchedule, exists
}

func (r *InMemoryFeedingScheduleRepo) GetAllByAnimalID(animalID uuid.UUID) []domain.FeedingSchedule {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var feedingSchedules []domain.FeedingSchedule
	for _, fs := range r.data {
		if fs.AnimalID() == animalID {
			feedingSchedules = append(feedingSchedules, *fs)
		}
	}
	return feedingSchedules
}
