package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"zoo/internal/application/dto"
	"zoo/internal/application/errs"
	"zoo/internal/application/port/in"
	"zoo/internal/presentation/response"
)

type FeedingHandler struct {
	feedingUseCase in.FeedingOrganizationUseCase
}

func NewFeedingHandler(feedingUseCase in.FeedingOrganizationUseCase) *FeedingHandler {
	return &FeedingHandler{feedingUseCase: feedingUseCase}
}

// GetAllSchedules godoc
// @Summary Get all feeding schedules
// @Tags feeding
// @Produce json
// @Success 200 {array} dto.FeedingScheduleResponse
// @Router /schedules [get]
func (h *FeedingHandler) GetAllSchedules(c *gin.Context) {
	schedules := h.feedingUseCase.GetAllFeedingSchedules()
	c.JSON(http.StatusOK, schedules)
}

// GetAnimalSchedules godoc
// @Summary Get feeding schedules for specific animal
// @Tags feeding
// @Produce json
// @Param id path string true "Animal ID"
// @Success 200 {array} dto.FeedingScheduleResponse
// @Failure 404 {object} map[string]string
// @Router /animals/{id}/schedules [get]
func (h *FeedingHandler) GetAnimalSchedules(c *gin.Context) {
	id := c.Param("id")

	schedules, err := h.feedingUseCase.GetAnimalFeedingSchedulesByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, schedules)
}

// NewSchedule godoc
// @Summary Create a new feeding schedule
// @Tags feeding
// @Accept json
// @Produce json
// @Param request body dto.NewFeedingScheduleRequest true "New feeding schedule"
// @Success 201 {object} dto.FeedingScheduleResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /schedules [post]
func (h *FeedingHandler) NewSchedule(c *gin.Context) {
	var req dto.NewFeedingScheduleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.HandleError(c, fmt.Errorf("%w: %v", errs.ErrInvalidRequest, err))
		return
	}

	schedule, err := h.feedingUseCase.NewFeedingSchedule(&req)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, schedule)
}

// FeedByScheduleId godoc
// @Summary Mark feeding schedule as occurred
// @Tags feeding
// @Produce json
// @Param id path string true "Feeding Schedule ID"
// @Success 200 {object} dto.FeedingScheduleResponse
// @Failure 404 {object} map[string]string
// @Router /schedules/{id}/feed [post]
func (h *FeedingHandler) FeedByScheduleId(c *gin.Context) {
	id := c.Param("id")

	schedule, err := h.feedingUseCase.FeedByScheduleId(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, schedule)
}
