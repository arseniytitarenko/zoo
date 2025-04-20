package in

type AnimalTransportUseCase interface {
	TransportAnimal(animalId, toEnclosureId string) error
}
