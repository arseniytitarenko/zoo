package dto

type AnimalStatisticsResponse struct {
	TotalCount  int                 `json:"total_count"`
	GenderStats GenderStatistics    `json:"gender_stats"`
	HealthStats HealthStatistics    `json:"health_stats"`
	Species     []SpeciesStatistics `json:"species_stats"`
}

type GenderStatistics struct {
	MaleCount   int `json:"male_count"`
	FemaleCount int `json:"female_count"`
}

type HealthStatistics struct {
	HealthyCount int `json:"healthy_count"`
	SickCount    int `json:"sick_count"`
}

type SpeciesStatistics struct {
	Species string `json:"species"`
	Count   int    `json:"count"`
}

type EnclosureStatisticsResponse struct {
	TotalCount                     int                     `json:"total_count"`
	TotalVolume                    float64                 `json:"total_volume"`
	TotalCapacity                  int                     `json:"total_capacity"`
	TotalAnimalCountInEnclosures   int                     `json:"total_animal_count_in_enclosures"`
	AverageAnimalCountPerEnclosure float64                 `json:"average_animal_count_per_enclosure"`
	AverageVolumePerAnimal         float64                 `json:"average_volume_per_animal"`
	TypeStats                      EnclosureTypeStatistics `json:"type_stats"`
}

type EnclosureTypeStatistics struct {
	AquariumCount      int `json:"aquarium_count"`
	ForPredatorsCount  int `json:"for_predators_count"`
	ForHerbivoresCount int `json:"for_herbivores_count"`
	ForBirdsCount      int `json:"for_birds_count"`
	OtherCount         int `json:"other_count"`
}

type FeedingStatisticsResponse struct {
	TotalCount    int    `json:"total_count"`
	OccurredCount int    `json:"occurred_count"`
	AverageDelay  string `json:"average_delay"`
}
