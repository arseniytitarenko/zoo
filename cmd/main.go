package main

import (
	"log"
	"zoo/internal/application/service"
	"zoo/internal/infrastructure/repository"
	"zoo/internal/presentation/handler"
	"zoo/internal/router"
)

func main() {
	animalRepo := repository.NewInMemoryAnimalRepo()
	animalService := service.NewAnimalService(animalRepo)
	animalHandler := handler.NewAnimalHandler(animalService)

	r := router.SetupRouter(animalHandler)
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
