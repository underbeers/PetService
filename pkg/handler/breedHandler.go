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

	if query.Has("petTypeID") {
		petTypeId, err := strconv.Atoi(query.Get("petTypeID"))
		if err != nil || petTypeId <= 0 {
			newErrorResponse(c, http.StatusBadRequest, "invalid pet type id")
			return
		}
		filter.PetTypeId = petTypeId
	}

	if query.Has("breedName") {
		filter.BreedName = query.Get("breedName")
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
