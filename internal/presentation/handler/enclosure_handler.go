package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"zoo/internal/application/dto"
	"zoo/internal/application/errs"
	"zoo/internal/application/port/in"
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
	if errors.Is(err, errs.ErrInvalidID) {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else if errors.Is(err, errs.ErrEnclosureNotFound) {
		c.JSON(http.StatusNotFound, err.Error())
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, enclosure)
}

func (h *EnclosureHandler) NewEnclosure(c *gin.Context) {
	var req dto.NewEnclosureRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	enclosure, err := h.enclosureUseCase.NewEnclosure(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, enclosure)
}

func (h *EnclosureHandler) DeleteEnclosure(c *gin.Context) {
	id := c.Param("id")

	err := h.enclosureUseCase.DeleteEnclosure(id)
	if errors.Is(err, errs.ErrInvalidID) {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else if errors.Is(err, errs.ErrEnclosureNotFound) {
		c.JSON(http.StatusNotFound, err.Error())
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
