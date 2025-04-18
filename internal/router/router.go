package router

import (
	"github.com/gin-gonic/gin"
	"zoo/internal/presentation/handler"
)

func SetupRouter(animalHandler *handler.AnimalHandler) *gin.Engine {
	r := gin.Default()

	animalGroup := r.Group("/animals")
	{
		animalGroup.GET("", animalHandler.GetAllAnimals)
		animalGroup.GET("/:id", animalHandler.GetAnimalByID)
		animalGroup.POST("", animalHandler.NewAnimal)
		animalGroup.DELETE("/:id", animalHandler.DeleteAnimal)
	}

	return r
}
