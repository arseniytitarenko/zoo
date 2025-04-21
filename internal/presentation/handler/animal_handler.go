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

// GetAllAnimals godoc
// @Summary Get all animals
// @Tags animals
// @Produce json
// @Success 200 {array} dto.AnimalResponse
// @Router /animals [get]
func (h *AnimalHandler) GetAllAnimals(c *gin.Context) {
	animals := h.animalUseCase.GetAllAnimals()
	c.JSON(http.StatusOK, animals)
}

// GetAnimalByID godoc
// @Summary Get animal by ID
// @Tags animals
// @Produce json
// @Param id path string true "Animal ID"
// @Success 200 {object} dto.AnimalResponse
// @Failure 404 {object} map[string]string
// @Router /animals/{id} [get]
func (h *AnimalHandler) GetAnimalByID(c *gin.Context) {
	id := c.Param("id")

	animal, err := h.animalUseCase.GetAnimalByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, animal)
}

// NewAnimal godoc
// @Summary Create a new animal
// @Tags animals
// @Accept json
// @Produce json
// @Param request body dto.NewAnimalRequest true "New animal data"
// @Success 201 {object} dto.AnimalResponse
// @Failure 400 {object} map[string]string
// @Router /animals [post]
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

// DeleteAnimal godoc
// @Summary Delete animal by ID
// @Tags animals
// @Produce json
// @Param id path string true "Animal ID"
// @Success 204 {object} nil
// @Failure 404 {object} map[string]string
// @Router /animals/{id} [delete]
func (h *AnimalHandler) DeleteAnimal(c *gin.Context) {
	id := c.Param("id")

	err := h.animalUseCase.DeleteAnimal(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

// TransportAnimal godoc
// @Summary Transport animal to another enclosure
// @Tags animals
// @Accept json
// @Produce json
// @Param id path string true "Animal ID"
// @Param request body dto.TransportAnimalRequest true "Target enclosure ID"
// @Success 204 {object} nil
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Router /animals/{id}/transport [post]
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
