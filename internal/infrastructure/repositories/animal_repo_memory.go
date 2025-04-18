package repositories

import (
	"github.com/google/uuid"
	"sync"
	"zoo/internal/domain"
)

type InMemoryAnimalRepo struct {
	data map[uuid.UUID]*domain.Animal
	mu   sync.RWMutex
}

func NewInMemoryAnimalRepo() *InMemoryAnimalRepo {
	return &InMemoryAnimalRepo{
		data: make(map[uuid.UUID]*domain.Animal),
	}
}

func (r *InMemoryAnimalRepo) Save(animal *domain.Animal) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[animal.ID()] = animal
}

func (r *InMemoryAnimalRepo) Delete(id uuid.UUID) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.data, id)
}

func (r *InMemoryAnimalRepo) GetAll() []domain.Animal {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var animals []domain.Animal
	for _, a := range r.data {
		animals = append(animals, *a)
	}
	return animals
}

func (r *InMemoryAnimalRepo) GetByID(id uuid.UUID) (*domain.Animal, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	animal, exists := r.data[id]
	return animal, exists
}
