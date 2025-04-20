package service

import (
	"github.com/google/uuid"
	"time"
	"zoo/internal/application/dto"
	"zoo/internal/application/errs"
	"zoo/internal/application/port/out"
	"zoo/internal/domain"
)

type AnimalService struct {
	animalRepo    out.AnimalRepository
	enclosureRepo out.EnclosureRepository
}

func NewAnimalService(animalRepo out.AnimalRepository) *AnimalService {
	return &AnimalService{animalRepo: animalRepo}
}

func (s *AnimalService) GetAllAnimals() []dto.AnimalResponse {
	animals := s.animalRepo.GetAll()
	return newAnimalResponses(animals)
}

func (s *AnimalService) GetAnimalByID(id string) (*dto.AnimalResponse, error) {
	animalId, err := uuid.Parse(id)
	if err != nil {
		return nil, errs.ErrInvalidID
	}
	animal, ok := s.animalRepo.GetByID(animalId)
	if !ok {
		return nil, errs.ErrAnimalNotFound
	}
	return newAnimalResponse(animal), nil
}

func (s *AnimalService) NewAnimal(req *dto.NewAnimalRequest) (*dto.AnimalResponse, error) {
	animalBirthDate, err := time.Parse("02.01.2006", req.BirthDate)
	if err != nil {
		return nil, errs.ErrInvalidDate
	}
	if req.Gender != domain.Male && req.Gender != domain.Female {
		return nil, errs.ErrInvalidGender
	}
	if req.HealthStatus != domain.Healthy && req.HealthStatus != domain.Sick {
		return nil, errs.ErrInvalidStatus
	}
	enclosure, ok := s.enclosureRepo.GetByID(req.EnclosureId)
	if !ok {
		return nil, errs.ErrEnclosureNotFound
	}
	if enclosure.IsFull() {
		return nil, errs.ErrEnclosureIsFull
	}
	_ = enclosure.AddAnimal()
	animal := domain.NewAnimal(
		req.Species,
		req.Name,
		animalBirthDate,
		req.Gender,
		req.FavoriteFood,
		req.HealthStatus,
		enclosure.ID(),
	)
	s.animalRepo.Save(animal)
	return newAnimalResponse(animal), nil
}

func (s *AnimalService) DeleteAnimal(id string) error {
	animalId, err := uuid.Parse(id)
	if err != nil {
		return errs.ErrInvalidID
	}
	animal, ok := s.animalRepo.GetByID(animalId)
	if !ok {
		return errs.ErrAnimalNotFound
	}
	s.animalRepo.Delete(animal.ID())
	return nil
}

func newAnimalResponses(animals []domain.Animal) []dto.AnimalResponse {
	result := make([]dto.AnimalResponse, len(animals))
	for i, animal := range animals {
		result[i] = *newAnimalResponse(&animal)
	}
	return result
}

func newAnimalResponse(animal *domain.Animal) *dto.AnimalResponse {
	return &dto.AnimalResponse{
		ID:           animal.ID(),
		EnclosureID:  animal.EnclosureID(),
		Species:      animal.Species(),
		Name:         animal.Name(),
		BirthDate:    animal.BirthDate().Format("02.01.2006"),
		Gender:       animal.Gender(),
		FavoriteFood: animal.FavoriteFood(),
		HealthStatus: animal.HealthStatus(),
	}
}
