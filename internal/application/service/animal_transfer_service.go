package service

import (
	"github.com/google/uuid"
	"zoo/internal/application/errs"
	"zoo/internal/application/port/out"
	"zoo/internal/domain"
)

type AnimalTransportService struct {
	animalRepo      out.AnimalRepository
	enclosureRepo   out.EnclosureRepository
	eventDispatcher out.IEventDispatcher
}

func NewAnimalTransportService(animalRepo out.AnimalRepository, enclosureRepo out.EnclosureRepository, dispatcher out.IEventDispatcher) *AnimalTransportService {
	return &AnimalTransportService{animalRepo: animalRepo, enclosureRepo: enclosureRepo, eventDispatcher: dispatcher}
}

func (s *AnimalTransportService) TransportAnimal(animalId, toEnclosureId string) error {
	animal, fromEnclosure, toEnclosure, err := s.GetAnimalAndEnclosures(animalId, toEnclosureId)
	if err != nil {
		return err
	}
	if toEnclosure.IsFull() {
		return errs.ErrEnclosureIsFull
	}
	err = fromEnclosure.RemoveAnimal()
	if err != nil {
		panic("enclosure remove failed: " + err.Error())
	}
	_ = toEnclosure.AddAnimal()
	animal.MoveTo(toEnclosure.ID)

	events := animal.PullEvents()
	for _, event := range events {
		s.eventDispatcher.Dispatch(event)
	}
	return nil
}

func (s *AnimalTransportService) GetAnimalAndEnclosures(animalId string, toEnclosureId string) (*domain.Animal,
	*domain.Enclosure, *domain.Enclosure, error) {
	animalUUID, err := uuid.Parse(animalId)
	if err != nil {
		return nil, nil, nil, errs.ErrInvalidID
	}
	toEnclosureUUID, err := uuid.Parse(toEnclosureId)
	if err != nil {
		return nil, nil, nil, errs.ErrInvalidID
	}
	animal, ok := s.animalRepo.GetByID(animalUUID)
	if !ok {
		return nil, nil, nil, errs.ErrAnimalNotFound
	}
	toEnclosure, ok := s.enclosureRepo.GetByID(toEnclosureUUID)
	if !ok {
		return nil, nil, nil, errs.ErrEnclosureNotFound
	}
	fromEnclosure, ok := s.enclosureRepo.GetByID(animal.EnclosureId)
	if !ok {
		panic("failed to find animal enclosure")
	}
	return animal, fromEnclosure, toEnclosure, nil
}
