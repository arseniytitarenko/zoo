package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"zoo/internal/application/dto"
	"zoo/internal/application/errs"
	"zoo/internal/application/port/in"
)

type AnimalHandler struct {
	animalUseCase          in.AnimalUseCase
	animalTransportUseCase in.AnimalTransportUseCase
}

func NewAnimalHandler(animalUseCase in.AnimalUseCase, animalTransportUseCase in.AnimalTransportUseCase) *AnimalHandler {
	return &AnimalHandler{animalUseCase: animalUseCase, animalTransportUseCase: animalTransportUseCase}
}

func (h *AnimalHandler) GetAllAnimals(c *gin.Context) {
	animals := h.animalUseCase.GetAllAnimals()
	c.JSON(http.StatusOK, animals)
}

func (h *AnimalHandler) GetAnimalByID(c *gin.Context) {
	id := c.Param("id")

	animal, err := h.animalUseCase.GetAnimalByID(id)
	if errors.Is(err, errs.ErrInvalidID) {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else if errors.Is(err, errs.ErrAnimalNotFound) {
		c.JSON(http.StatusNotFound, err.Error())
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, animal)
}

func (h *AnimalHandler) NewAnimal(c *gin.Context) {
	var req dto.NewAnimalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	animal, err := h.animalUseCase.NewAnimal(&req)
	if errors.Is(err, errs.ErrInvalidDate) || errors.Is(err, errs.ErrInvalidGender) ||
		errors.Is(err, errs.ErrInvalidStatus) || errors.Is(err, errs.ErrEnclosureNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if errors.Is(err, errs.ErrEnclosureIsFull) {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, animal)
}

func (h *AnimalHandler) DeleteAnimal(c *gin.Context) {
	id := c.Param("id")

	err := h.animalUseCase.DeleteAnimal(id)
	if errors.Is(err, errs.ErrInvalidID) {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else if errors.Is(err, errs.ErrAnimalNotFound) {
		c.JSON(http.StatusNotFound, err.Error())
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func (h *AnimalHandler) TransportAnimal(c *gin.Context) {
	id := c.Param("id")
	var req dto.TransportAnimalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.animalTransportUseCase.TransportAnimal(id, req.ToEnclosureID)
	if errors.Is(err, errs.ErrEnclosureNotFound) || errors.Is(err, errs.ErrInvalidID) {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else if errors.Is(err, errs.ErrAnimalNotFound) {
		c.JSON(http.StatusNotFound, err.Error())
		return
	} else if errors.Is(err, errs.ErrEnclosureIsFull) {
		c.JSON(http.StatusConflict, err.Error())
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
