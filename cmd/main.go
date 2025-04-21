package main

import (
	"log"
	"zoo/internal/application/service"
	"zoo/internal/infrastructure/dispatcher"
	"zoo/internal/infrastructure/repository"
	_ "zoo/internal/presentation/docs"
	"zoo/internal/presentation/handler"
	"zoo/internal/presentation/router"
)

// @title Zoo API
// @version 1.0
// @description API для зоопарка
// @host localhost:8080
// @BasePath

func main() {
	animalRepo := repository.NewInMemoryAnimalRepo()
	enclosureRepo := repository.NewInMemoryEnclosureRepo()
	feedingRepo := repository.NewInMemoryFeedingScheduleRepo()
	eventDispatcher := dispatcher.NewEventDispatcher()
	dispatcher.RegisterEventHandlers(eventDispatcher)

	animalService := service.NewAnimalService(animalRepo, enclosureRepo)
	enclosureService := service.NewEnclosureService(enclosureRepo)
	animalTransportService := service.NewAnimalTransportService(animalRepo, enclosureRepo, eventDispatcher)
	feedingService := service.NewFeedingOrganizationService(animalRepo, feedingRepo, eventDispatcher)
	zooStatisticsService := service.NewZooStatisticsService(animalRepo, feedingRepo, enclosureRepo)

	animalHandler := handler.NewAnimalHandler(animalService, animalTransportService)
	enclosureHandler := handler.NewEnclosureHandler(enclosureService)
	feedingHandler := handler.NewFeedingHandler(feedingService)
	zooStatisticsHandler := handler.NewZooStatisticsHandler(zooStatisticsService)

	r := router.SetupRouter(animalHandler, enclosureHandler, feedingHandler, zooStatisticsHandler)
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
