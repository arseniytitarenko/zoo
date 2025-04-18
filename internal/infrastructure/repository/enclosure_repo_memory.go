package repository

import (
	"github.com/google/uuid"
	"sync"
	"zoo/internal/domain"
)

type InMemoryEnclosureRepo struct {
	data map[uuid.UUID]*domain.Enclosure
	mu   sync.RWMutex
}

func NewInMemoryEnclosureRepo() *InMemoryEnclosureRepo {
	return &InMemoryEnclosureRepo{
		data: make(map[uuid.UUID]*domain.Enclosure),
	}
}

func (r *InMemoryEnclosureRepo) Save(enclosure *domain.Enclosure) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[enclosure.ID()] = enclosure
}

func (r *InMemoryEnclosureRepo) Delete(id uuid.UUID) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.data, id)
}

func (r *InMemoryEnclosureRepo) GetAll() []domain.Enclosure {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var enclosures []domain.Enclosure
	for _, a := range r.data {
		enclosures = append(enclosures, *a)
	}
	return enclosures
}

func (r *InMemoryEnclosureRepo) GetByID(id uuid.UUID) (*domain.Enclosure, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	enclosure, exists := r.data[id]
	return enclosure, exists
}
