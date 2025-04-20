package router

import (
	"github.com/gin-gonic/gin"
	"zoo/internal/presentation/handler"
)

func SetupRouter(animalHandler *handler.AnimalHandler, enclosureHandler *handler.EnclosureHandler) *gin.Engine {
	r := gin.Default()

	animalGroup := r.Group("/animals")
	{
		animalGroup.GET("", animalHandler.GetAllAnimals)
		animalGroup.GET("/:id", animalHandler.GetAnimalByID)
		animalGroup.POST("", animalHandler.NewAnimal)
		animalGroup.DELETE("/:id", animalHandler.DeleteAnimal)
	}

	enclosureGroup := r.Group("/enclosures")
	{
		enclosureGroup.GET("", enclosureHandler.GetAllEnclosures)
		enclosureGroup.GET("/:id", enclosureHandler.GetEnclosureByID)
		enclosureGroup.POST("", enclosureHandler.NewEnclosure)
		enclosureGroup.DELETE("/:id", enclosureHandler.DeleteEnclosure)
	}

	return r
}
