package service

import (
	"github.com/google/uuid"
	"zoo/internal/application/dto"
	"zoo/internal/application/errs"
	"zoo/internal/application/port/out"
	"zoo/internal/domain"
)

type EnclosureService struct {
	enclosureRepo out.EnclosureRepository
}

func NewEnclosureService(enclosureRepo out.EnclosureRepository) *EnclosureService {
	return &EnclosureService{enclosureRepo: enclosureRepo}
}

func (s *EnclosureService) GetAllEnclosures() []dto.EnclosureResponse {
	animals := s.enclosureRepo.GetAll()
	return newEnclosureResponses(animals)
}

func (s *EnclosureService) GetEnclosureByID(id string) (*dto.EnclosureResponse, error) {
	enclosureId, err := uuid.Parse(id)
	if err != nil {
		return nil, errs.ErrInvalidID
	}
	animal, ok := s.enclosureRepo.GetByID(enclosureId)
	if !ok {
		return nil, errs.ErrEnclosureNotFound
	}
	return newEnclosureResponse(animal), nil
}

func (s *EnclosureService) NewEnclosure(req *dto.NewEnclosureRequest) (*dto.EnclosureResponse, error) {
	if req.Type != domain.Aquarium && req.Type != domain.ForBirds &&
		req.Type != domain.ForHerbivores && req.Type != domain.ForPredators {
		req.Type = domain.Other
	}

	enclosure := domain.NewEnclosure(req.Type, domain.NewSize(req.Length, req.Width, req.Height), req.AnimalCapacity)

	s.enclosureRepo.Save(enclosure)
	return newEnclosureResponse(enclosure), nil
}

func (s *EnclosureService) DeleteEnclosure(id string) error {
	enclosureID, err := uuid.Parse(id)
	if err != nil {
		return errs.ErrInvalidID
	}
	enclosure, ok := s.enclosureRepo.GetByID(enclosureID)
	if !ok {
		return errs.ErrEnclosureNotFound
	}
	s.enclosureRepo.Delete(enclosure.ID())
	return nil
}

func newEnclosureResponses(enclosures []domain.Enclosure) []dto.EnclosureResponse {
	result := make([]dto.EnclosureResponse, len(enclosures))
	for i, enclosure := range enclosures {
		result[i] = *newEnclosureResponse(&enclosure)
	}
	return result
}

func newEnclosureResponse(enclosure *domain.Enclosure) *dto.EnclosureResponse {
	return &dto.EnclosureResponse{
		ID:              enclosure.ID(),
		Length:          enclosure.Size().Length(),
		Width:           enclosure.Size().Width(),
		Height:          enclosure.Size().Height(),
		Type:            enclosure.Type(),
		CurrAnimalCount: enclosure.CurrAnimalCount(),
		AnimalCapacity:  enclosure.AnimalCapacity(),
	}
}
