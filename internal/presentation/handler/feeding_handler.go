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

func (h *FeedingHandler) GetAllSchedules(c *gin.Context) {
	schedules := h.feedingUseCase.GetAllFeedingSchedules()
	c.JSON(http.StatusOK, schedules)
}

func (h *FeedingHandler) GetAnimalSchedules(c *gin.Context) {
	id := c.Param("id")

	schedules, err := h.feedingUseCase.GetAnimalFeedingSchedulesByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, schedules)
}

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

func (h *FeedingHandler) FeedByScheduleId(c *gin.Context) {
	id := c.Param("id")

	schedule, err := h.feedingUseCase.FeedByScheduleId(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, schedule)
}
