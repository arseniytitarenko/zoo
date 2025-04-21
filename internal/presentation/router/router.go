package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"zoo/internal/presentation/handler"
)

func SetupRouter(
	animalHandler *handler.AnimalHandler,
	enclosureHandler *handler.EnclosureHandler,
	feedingHandler *handler.FeedingHandler,
	zooStatisticsHandler *handler.ZooStatisticsHandler,
) *gin.Engine {

	r := gin.Default()

	animalGroup := r.Group("/animals")
	{
		animalGroup.GET("", animalHandler.GetAllAnimals)
		animalGroup.POST("", animalHandler.NewAnimal)
		animalGroup.GET("/:id", animalHandler.GetAnimalByID)
		animalGroup.DELETE("/:id", animalHandler.DeleteAnimal)
		animalGroup.POST("/:id/transport", animalHandler.TransportAnimal)
		animalGroup.GET("/:id/schedules", feedingHandler.GetAnimalSchedules)
	}

	enclosureGroup := r.Group("/enclosures")
	{
		enclosureGroup.GET("", enclosureHandler.GetAllEnclosures)
		enclosureGroup.GET("/:id", enclosureHandler.GetEnclosureByID)
		enclosureGroup.POST("", enclosureHandler.NewEnclosure)
		enclosureGroup.DELETE("/:id", enclosureHandler.DeleteEnclosure)
	}

	scheduleGroup := r.Group("/schedules")
	{
		scheduleGroup.GET("", feedingHandler.GetAllSchedules)
		scheduleGroup.POST("", feedingHandler.NewSchedule)
		scheduleGroup.POST("/:id/feed", feedingHandler.FeedByScheduleId)
	}

	statisticsGroup := r.Group("/statistics")
	{
		statisticsGroup.GET("/animals", zooStatisticsHandler.GetAnimalStatistics)
		statisticsGroup.GET("/enclosures", zooStatisticsHandler.GetEnclosureStatistics)
		statisticsGroup.GET("/schedules", zooStatisticsHandler.GetFeedingStatistics)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
