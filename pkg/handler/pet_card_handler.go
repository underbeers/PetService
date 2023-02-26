package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/underbeers/PetService/pkg/models"
	"net/http"
	"strconv"
)

func (h *Handler) createNewCard(c *gin.Context) {

	var input models.PetCard
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if _, err := h.services.PetType.GetAll(models.PetTypeFilter{PetTypeId: input.PetTypeId}); err != nil {
		c.JSON(http.StatusBadRequest, statusResponse{"incorrect pet type id"})
		return
	}

	breedId, err := h.services.Breed.GetAll(models.BreedFilter{BreedId: input.BreedId})
	if err != nil || breedId[0].PetTypeId != input.PetTypeId {
		c.JSON(http.StatusBadRequest, statusResponse{"incorrect breed id"})
		return
	}

	err = h.services.PetCard.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) getAllCards(c *gin.Context) {
	query := c.Request.URL.Query()
	filter := models.PetCardFilter{}

	if query.Has("id") {
		PetCardId, err := strconv.Atoi(query.Get("id"))
		if err != nil || PetCardId <= 0 {
			newErrorResponse(c, http.StatusBadRequest, "invalid pet card id param")
			return
		}
		filter.PetCardId = PetCardId
	}

	if query.Has("user_id") {
		UserId, err := strconv.Atoi(query.Get("user_id"))
		if err != nil || UserId <= 0 {
			newErrorResponse(c, http.StatusBadRequest, "invalid user id param")
			return
		}
		filter.UserId = UserId
	}

	petCardList, err := h.services.PetCard.GetAll(filter)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(petCardList) == 0 {
		newErrorResponse(c, http.StatusBadRequest, "records not found")
		return
	}

	c.JSON(http.StatusOK, petCardList)
}

func (h *Handler) getMainCardInfo(c *gin.Context) {
	query := c.Request.URL.Query()
	filter := models.PetCardFilter{}

	if query.Has("id") {
		PetCardId, err := strconv.Atoi(query.Get("id"))
		if err != nil || PetCardId <= 0 {
			newErrorResponse(c, http.StatusBadRequest, "invalid pet card id param")
			return
		}
		filter.PetCardId = PetCardId
	}

	if query.Has("user_id") {
		UserId, err := strconv.Atoi(query.Get("user_id"))
		if err != nil || UserId <= 0 {
			newErrorResponse(c, http.StatusBadRequest, "invalid user id param")
			return
		}
		filter.UserId = UserId
	}

	petCardList, err := h.services.PetCard.GetMain(filter)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(petCardList) == 0 {
		newErrorResponse(c, http.StatusBadRequest, "records not found")
		return
	}

	c.JSON(http.StatusOK, petCardList)
}

func (h *Handler) updateCard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	petCard, err := h.services.PetCard.GetAll(models.PetCardFilter{PetCardId: id})
	if len(petCard) != 1 || err != nil {
		c.JSON(http.StatusBadRequest, statusResponse{"incorrect pet card id"})
		return
	}

	var input models.UpdateCardInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.PetTypeId != nil {
		if _, err := h.services.PetType.GetAll(models.PetTypeFilter{PetTypeId: *input.PetTypeId}); err != nil {
			c.JSON(http.StatusBadRequest, statusResponse{"incorrect pet type id"})
			return
		}
	}

	if input.BreedId != nil {
		breedId, err := h.services.Breed.GetAll(models.BreedFilter{BreedId: *input.BreedId})
		if err != nil || breedId[0].PetTypeId != *input.PetTypeId {
			c.JSON(http.StatusBadRequest, statusResponse{"incorrect breed id"})
			return
		}
	}

	if err := h.services.PetCard.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteCard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	petCard, err := h.services.PetCard.GetAll(models.PetCardFilter{PetCardId: id})
	if len(petCard) != 1 || err != nil {
		c.JSON(http.StatusBadRequest, statusResponse{"incorrect pet card id"})
		return
	}

	err = h.services.PetCard.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}