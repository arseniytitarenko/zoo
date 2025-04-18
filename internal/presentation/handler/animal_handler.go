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
	animalUseCase in.AnimalUseCase
}

func NewAnimalHandler(animalUseCase in.AnimalUseCase) *AnimalHandler {
	return &AnimalHandler{animalUseCase: animalUseCase}
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
