package domain

import (
	"errors"
	"github.com/google/uuid"
)

type Enclosure struct {
	id              uuid.UUID
	enclosureType   EnclosureType
	size            *Size
	currAnimalCount uint8
	animalCapacity  uint8
}

func NewEnclosure(enclosureType EnclosureType, size *Size, capacity uint8) *Enclosure {
	return &Enclosure{
		id:             uuid.New(),
		enclosureType:  enclosureType,
		size:           size,
		animalCapacity: capacity,
	}
}

func (e *Enclosure) Clean() {}

func (e *Enclosure) AddAnimal() error {
	if e.currAnimalCount == e.animalCapacity {
		return errors.New("enclosure is full")
	}
	e.currAnimalCount++
	return nil
}

func (e *Enclosure) RemoveAnimal() error {
	if e.currAnimalCount == 0 {
		return errors.New("enclosure is empty")
	}
	e.currAnimalCount--
	return nil
}

func (e *Enclosure) IsFull() bool {
	return e.currAnimalCount == e.animalCapacity
}

func (e *Enclosure) ID() uuid.UUID          { return e.id }
func (e *Enclosure) Type() EnclosureType    { return e.enclosureType }
func (e *Enclosure) Size() *Size            { return e.size }
func (e *Enclosure) CurrAnimalCount() uint8 { return e.currAnimalCount }
func (e *Enclosure) AnimalCapacity() uint8  { return e.animalCapacity }

type EnclosureType string

const (
	Aquarium      EnclosureType = "Aquarium"
	ForPredators  EnclosureType = "ForPredators"
	ForHerbivores EnclosureType = "ForHerbivores"
	ForBirds      EnclosureType = "ForBirds"
	Other         EnclosureType = "Other"
)

// --- Size ---

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

func (s *Size) Length() float64 { return s.length }
func (s *Size) Width() float64  { return s.width }
func (s *Size) Height() float64 { return s.height }
func (s *Size) Volume() float64 { return s.length * s.width * s.height }
