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

// GetAllEnclosures godoc
// @Summary Get all enclosures
// @Tags enclosures
// @Produce json
// @Success 200 {array} dto.EnclosureResponse
// @Router /enclosures [get]
func (h *EnclosureHandler) GetAllEnclosures(c *gin.Context) {
	enclosures := h.enclosureUseCase.GetAllEnclosures()
	c.JSON(http.StatusOK, enclosures)
}

// GetEnclosureByID godoc
// @Summary Get enclosure by ID
// @Tags enclosures
// @Produce json
// @Param id path string true "Enclosure ID"
// @Success 200 {object} dto.EnclosureResponse
// @Failure 404 {object} map[string]string
// @Router /enclosures/{id} [get]
func (h *EnclosureHandler) GetEnclosureByID(c *gin.Context) {
	id := c.Param("id")

	enclosure, err := h.enclosureUseCase.GetEnclosureByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, enclosure)
}

// NewEnclosure godoc
// @Summary Create a new enclosure
// @Tags enclosures
// @Accept json
// @Produce json
// @Param request body dto.NewEnclosureRequest true "New enclosure data"
// @Success 201 {object} dto.EnclosureResponse
// @Failure 400 {object} map[string]string
// @Router /enclosures [post]
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

// DeleteEnclosure godoc
// @Summary Delete enclosure by ID
// @Tags enclosures
// @Produce json
// @Param id path string true "Enclosure ID"
// @Success 204 {object} nil
// @Failure 404 {object} map[string]string
// @Router /enclosures/{id} [delete]
func (h *EnclosureHandler) DeleteEnclosure(c *gin.Context) {
	id := c.Param("id")

	err := h.enclosureUseCase.DeleteEnclosure(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
