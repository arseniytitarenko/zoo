package repository

import (
	"github.com/google/uuid"
	"zoo/internal/domain"
)

type IFeedingScheduleRepository interface {
	Save(animal *domain.FeedingSchedule)
	Delete(id uuid.UUID)
	GetByID(id uuid.UUID) (*domain.FeedingSchedule, bool)
	GetAll() []domain.FeedingSchedule
}
