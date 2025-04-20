package main

import (
	"log"
	"zoo/internal/application/service"
	"zoo/internal/infrastructure/dispatcher"
	"zoo/internal/infrastructure/repository"
	"zoo/internal/presentation/handler"
	"zoo/internal/presentation/router"
)

func main() {
	animalRepo := repository.NewInMemoryAnimalRepo()
	enclosureRepo := repository.NewInMemoryEnclosureRepo()
	feedingRepo := repository.NewInMemoryFeedingScheduleRepo()
	eventDispatcher := dispatcher.NewEventDispatcher()
	dispatcher.RegisterEventHandlers(eventDispatcher)

	animalService := service.NewAnimalService(animalRepo)
	enclosureService := service.NewEnclosureService(enclosureRepo)
	animalTransportService := service.NewAnimalTransportService(animalRepo, enclosureRepo, eventDispatcher)
	feedingService := service.NewFeedingOrganizationService(animalRepo, feedingRepo, eventDispatcher)

	animalHandler := handler.NewAnimalHandler(animalService, animalTransportService)
	enclosureHandler := handler.NewEnclosureHandler(enclosureService)
	feedingHandler := handler.NewFeedingHandler(feedingService)

	r := router.SetupRouter(animalHandler, enclosureHandler, feedingHandler)
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
