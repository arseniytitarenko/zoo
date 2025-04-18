package out

import (
	"github.com/google/uuid"
	"zoo/internal/domain"
)

type EnclosureRepository interface {
	Save(animal *domain.Enclosure)
	Delete(id uuid.UUID)
	GetByID(id uuid.UUID) (*domain.Enclosure, bool)
	GetAll() []domain.Enclosure
}
