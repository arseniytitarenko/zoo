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

// GetAnimalStatistics godoc
// @Summary Get animal statistics
// @Tags statistics
// @Produce json
// @Success 200 {object} dto.AnimalStatisticsResponse
// @Router /statistics/animals [get]
func (h *ZooStatisticsHandler) GetAnimalStatistics(c *gin.Context) {
	animalStatistics := h.zooStatistics.CalculateAnimalStatistics()
	c.JSON(http.StatusOK, animalStatistics)
}

// GetEnclosureStatistics godoc
// @Summary Get enclosure statistics
// @Tags statistics
// @Produce json
// @Success 200 {object} dto.EnclosureStatisticsResponse
// @Router /statistics/enclosures [get]
func (h *ZooStatisticsHandler) GetEnclosureStatistics(c *gin.Context) {
	enclosureStatistics := h.zooStatistics.CalculateEnclosureStatistics()
	c.JSON(http.StatusOK, enclosureStatistics)
}

// GetFeedingStatistics godoc
// @Summary Get feeding statistics
// @Tags statistics
// @Produce json
// @Success 200 {object} dto.FeedingStatisticsResponse
// @Router /statistics/schedules [get]
func (h *ZooStatisticsHandler) GetFeedingStatistics(c *gin.Context) {
	feedingStatistics := h.zooStatistics.CalculateFeedingStatistics()
	c.JSON(http.StatusOK, feedingStatistics)
}
