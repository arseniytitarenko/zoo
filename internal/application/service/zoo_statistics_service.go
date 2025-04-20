package service

import (
	"fmt"
	"time"
	"zoo/internal/application/dto"
	"zoo/internal/application/port/out"
	"zoo/internal/domain"
)

type ZooStatisticsService struct {
	animalRepo          out.AnimalRepository
	enclosureRepo       out.EnclosureRepository
	feedingScheduleRepo out.FeedingScheduleRepository
}

func NewZooStatisticsService(
	animalRepo out.AnimalRepository,
	feedingScheduleRepo out.FeedingScheduleRepository,
	enclosureRepo out.EnclosureRepository) *ZooStatisticsService {
	return &ZooStatisticsService{
		animalRepo:          animalRepo,
		feedingScheduleRepo: feedingScheduleRepo,
		enclosureRepo:       enclosureRepo,
	}
}

func (s *ZooStatisticsService) CalculateAnimalStatistics() *dto.AnimalStatisticsResponse {
	animals := s.animalRepo.GetAll()
	totalCount := len(animals)
	maleCount, femaleCount, healthyCount, sickCount := 0, 0, 0, 0
	speciesStats := make(map[string]int, totalCount)
	for _, animal := range animals {
		if animal.Gender() == domain.Male {
			maleCount++
		}
		if animal.Gender() == domain.Female {
			femaleCount++
		}
		if animal.HealthStatus() == domain.Healthy {
			healthyCount++
		}
		if animal.HealthStatus() == domain.Sick {
			sickCount++
		}
		speciesStats[animal.Species()]++
	}
	var speciesStatsSlice []dto.SpeciesStatistics
	for species, count := range speciesStats {
		speciesStatsSlice = append(speciesStatsSlice, dto.SpeciesStatistics{
			Species: species,
			Count:   count,
		})
	}
	return &dto.AnimalStatisticsResponse{
		TotalCount: totalCount,
		Species:    speciesStatsSlice,
		GenderStats: dto.GenderStatistics{
			MaleCount:   maleCount,
			FemaleCount: femaleCount,
		},
		HealthStats: dto.HealthStatistics{
			HealthyCount: healthyCount,
			SickCount:    sickCount,
		},
	}

}

func (s *ZooStatisticsService) CalculateEnclosureStatistics() *dto.EnclosureStatisticsResponse {
	enclosures := s.enclosureRepo.GetAll()
	totalCount := len(enclosures)
	var volume float64
	var animalCount, capacity int
	aquariumCount, forPredatorsCount, forHerbivoresCount, forBirdsCount, otherCount := 0, 0, 0, 0, 0
	for _, enclosure := range enclosures {
		volume += enclosure.Size().Volume()
		animalCount += int(enclosure.CurrAnimalCount())
		capacity += int(enclosure.AnimalCapacity())
		switch enclosure.Type() {
		case domain.Aquarium:
			aquariumCount++
		case domain.ForBirds:
			forBirdsCount++
		case domain.ForHerbivores:
			forHerbivoresCount++
		case domain.ForPredators:
			forPredatorsCount++
		default:
			otherCount++
		}
	}
	return &dto.EnclosureStatisticsResponse{
		TotalCount:                     totalCount,
		TotalVolume:                    volume,
		TotalCapacity:                  capacity,
		TotalAnimalCountInEnclosures:   animalCount,
		AverageAnimalCountPerEnclosure: float64(animalCount) / float64(totalCount),
		AverageVolumePerAnimal:         volume / float64(totalCount),
		TypeStats: dto.EnclosureTypeStatistics{
			AquariumCount:      aquariumCount,
			ForBirdsCount:      forBirdsCount,
			ForHerbivoresCount: forHerbivoresCount,
			ForPredatorsCount:  forPredatorsCount,
			OtherCount:         otherCount,
		},
	}
}

func (s *ZooStatisticsService) CalculateFeedingStatistics() *dto.FeedingStatisticsResponse {
	feedings := s.feedingScheduleRepo.GetAll()
	totalCount := len(feedings)
	occurredCount := 0
	var averageDelay time.Duration
	for _, feeding := range feedings {
		if feeding.IsOccurred() {
			occurredCount++
		}
		averageDelay += feeding.OccurredAt().Sub(feeding.ScheduledAt())
	}
	return &dto.FeedingStatisticsResponse{
		TotalCount:    totalCount,
		OccurredCount: occurredCount,
		AverageDelay:  FormatDuration(averageDelay / time.Duration(totalCount)),
	}
}

func FormatDuration(d time.Duration) string {
	totalSeconds := int(d.Seconds())
	hours := totalSeconds / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60
	return fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
}
