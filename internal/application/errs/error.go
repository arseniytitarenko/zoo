package errs

import "errors"

var (
	ErrAnimalNotFound    = errors.New("animal not found")
	ErrEnclosureNotFound = errors.New("enclosure not found")
	ErrEnclosureIsFull   = errors.New("impossible to place the animal, the cage is full")
	ErrInvalidID         = errors.New("invalid id")
	ErrInvalidDate       = errors.New("invalid date")
	ErrInvalidGender     = errors.New("invalid gender: gender should be one of: Male, Female")
	ErrInvalidStatus     = errors.New("invalid health status: should be one of: Healthy, Sick")
)
