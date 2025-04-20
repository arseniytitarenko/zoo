package in

import "zoo/internal/application/dto"

type ZooStatisticsUseCase interface {
	CalculateAnimalStatistics() *dto.AnimalStatisticsResponse
	CalculateEnclosureStatistics() *dto.EnclosureStatisticsResponse
	CalculateFeedingStatistics() *dto.FeedingStatisticsResponse
}
