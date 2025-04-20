package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zoo/internal/application/port/in"
)

type ZooStatisticsHandler struct {
	zooStatistics in.ZooStatisticsUseCase
}

func NewZooStatisticsHandler(zooStatistics in.ZooStatisticsUseCase) *ZooStatisticsHandler {
	return &ZooStatisticsHandler{zooStatistics: zooStatistics}
}

func (h *ZooStatisticsHandler) GetAnimalStatistics(c *gin.Context) {
	animalStatistics := h.zooStatistics.CalculateAnimalStatistics()
	c.JSON(http.StatusOK, animalStatistics)
}

func (h *ZooStatisticsHandler) GetEnclosureStatistics(c *gin.Context) {
	enclosureStatistics := h.zooStatistics.CalculateEnclosureStatistics()
	c.JSON(http.StatusOK, enclosureStatistics)
}

func (h *ZooStatisticsHandler) GetFeedingStatistics(c *gin.Context) {
	feedingStatistics := h.zooStatistics.CalculateFeedingStatistics()
	c.JSON(http.StatusOK, feedingStatistics)
}
