package handler

import (
	"github.com/gin-gonic/gin"
	pet_service "github.com/underbeers/PetService/pkg/models"
	"net/http"
	"strconv"
)

func (h *Handler) getAllTypes(c *gin.Context) {
	query := c.Request.URL.Query()
	filter := pet_service.PetTypeFilter{}

	if query.Has("id") {
		PetTypeId, err := strconv.Atoi(query.Get("id"))
		if err != nil || PetTypeId <= 0 {
			newErrorResponse(c, http.StatusBadRequest, "invalid pet type id")
			return
		}
		filter.PetTypeId = PetTypeId
	}

	if query.Has("pet_type") {
		filter.PetType = query.Get("pet_type")
	}

	petTypeList, err := h.services.PetType.GetAll(filter)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(petTypeList) == 0 {
		newErrorResponse(c, http.StatusBadRequest, "records not found")
		return
	}

	c.JSON(http.StatusOK, petTypeList)
}
