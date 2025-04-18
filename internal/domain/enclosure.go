package domain

import (
	"errors"
	"github.com/google/uuid"
)

type Enclosure struct {
	id              uuid.UUID
	Type            EnclosureType
	Size            Size
	CurrAnimalCount uint8
	AnimalCapacity  uint8
}

func (e *Enclosure) ID() uuid.UUID {
	return e.id
}

func (e *Enclosure) Clean() {}

func (e *Enclosure) AddAnimal() error {
	if e.CurrAnimalCount == e.AnimalCapacity {
		return errors.New("enclosure is full")
	}
	e.CurrAnimalCount++
	return nil
}

func (e *Enclosure) RemoveAnimal() error {
	if e.CurrAnimalCount == 0 {
		return errors.New("enclosure is empty")
	}
	e.CurrAnimalCount--
	return nil
}

type EnclosureType string

const (
	Aquarium      EnclosureType = "Aquarium"
	ForPredators  EnclosureType = "For predators"
	ForHerbivores EnclosureType = "For herbivores"
	ForBirds      EnclosureType = "For birds"
	Other         EnclosureType = "Other"
)

type Size struct {
	length float64
	width  float64
	height float64
}

func NewSize(length, width, height float64) *Size {
	return &Size{
		length: length,
		width:  width,
		height: height,
	}
}
