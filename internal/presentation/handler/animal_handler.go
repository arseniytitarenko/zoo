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
	if err != nil {
		response.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, animal)
}

func (h *AnimalHandler) NewAnimal(c *gin.Context) {
	var req dto.NewAnimalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.HandleError(c, fmt.Errorf("%w: %v", errs.ErrInvalidRequest, err))
		return
	}

	animal, err := h.animalUseCase.NewAnimal(&req)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, animal)
}

func (h *AnimalHandler) DeleteAnimal(c *gin.Context) {
	id := c.Param("id")

	err := h.animalUseCase.DeleteAnimal(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func (h *AnimalHandler) TransportAnimal(c *gin.Context) {
	id := c.Param("id")
	var req dto.TransportAnimalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.HandleError(c, fmt.Errorf("%w: %v", errs.ErrInvalidRequest, err))
		return
	}

	err := h.animalTransportUseCase.TransportAnimal(id, req.ToEnclosureID)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
