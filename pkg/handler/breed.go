package handler

import (
	"github.com/gin-gonic/gin"
	pet_service "github.com/underbeers/PetService/pkg/models"
	"net/http"
	"strconv"
)

func (h *Handler) getAllBreeds(c *gin.Context) {
	query := c.Request.URL.Query()
	filter := pet_service.BreedFilter{}

	if query.Has("id") {
		BreedId, err := strconv.Atoi(query.Get("id"))
		if err != nil || BreedId <= 0 {
			newErrorResponse(c, http.StatusBadRequest, "invalid breed id")
			return
		}
		filter.BreedId = BreedId
	}

	if query.Has("pet_type_id") {
		petTypeId, err := strconv.Atoi(query.Get("pet_type_id"))
		if err != nil || petTypeId <= 0 {
			newErrorResponse(c, http.StatusBadRequest, "invalid pet type id")
			return
		}
		filter.PetTypeId = petTypeId
	}

	if query.Has("breed_name") {
		filter.BreedName = query.Get("breed_name")
	}

	breedList, err := h.services.Breed.GetAll(filter)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(breedList) == 0 {
		newErrorResponse(c, http.StatusBadRequest, "records not found")
		return
	}

	c.JSON(http.StatusOK, breedList)
}
