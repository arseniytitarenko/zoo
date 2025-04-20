package out

import (
	"github.com/google/uuid"
	"zoo/internal/domain"
)

type FeedingScheduleRepository interface {
	Save(animal *domain.FeedingSchedule)
	Delete(id uuid.UUID)
	GetByID(id uuid.UUID) (*domain.FeedingSchedule, bool)
	GetAll() []domain.FeedingSchedule
	GetAllByAnimalID(animalID uuid.UUID) []domain.FeedingSchedule
}
