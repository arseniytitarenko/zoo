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

type EnclosureHandler struct {
	enclosureUseCase in.EnclosureUseCase
}

func NewEnclosureHandler(enclosureUseCase in.EnclosureUseCase) *EnclosureHandler {
	return &EnclosureHandler{enclosureUseCase: enclosureUseCase}
}

func (h *EnclosureHandler) GetAllEnclosures(c *gin.Context) {
	enclosures := h.enclosureUseCase.GetAllEnclosures()
	c.JSON(http.StatusOK, enclosures)
}

func (h *EnclosureHandler) GetEnclosureByID(c *gin.Context) {
	id := c.Param("id")

	enclosure, err := h.enclosureUseCase.GetEnclosureByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, enclosure)
}

func (h *EnclosureHandler) NewEnclosure(c *gin.Context) {
	var req dto.NewEnclosureRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.HandleError(c, fmt.Errorf("%w: %v", errs.ErrInvalidRequest, err))
		return
	}

	enclosure, err := h.enclosureUseCase.NewEnclosure(&req)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, enclosure)
}

func (h *EnclosureHandler) DeleteEnclosure(c *gin.Context) {
	id := c.Param("id")

	err := h.enclosureUseCase.DeleteEnclosure(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
