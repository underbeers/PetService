package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/underbeers/PetService/pkg/models"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (h *Handler) createNewCard(c *gin.Context) {

	var input models.PetCard
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	/*Проверка, что такой pet type id существует*/
	if _, err := h.services.PetType.GetAll(models.PetTypeFilter{PetTypeId: input.PetTypeId}); err != nil {
		c.JSON(http.StatusBadRequest, statusResponse{"incorrect pet type id"})
		return
	}

	/*Проверка, что такой breed id существует*/
	breedId, err := h.services.Breed.GetAll(models.BreedFilter{BreedId: input.BreedId})
	if err != nil || len(breedId) == 0 || breedId[0].PetTypeId != input.PetTypeId {
		c.JSON(http.StatusBadRequest, statusResponse{"incorrect breed id"})
		return
	}

	userID := c.Request.Header.Get("userID")

	if len(userID) == 0 {
		c.JSON(http.StatusBadRequest, statusResponse{"invalid access token"})
		return
	}
	id, err := uuid.Parse(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	input.UserId = id

	err = h.services.PetCard.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) getAllCards(c *gin.Context) {

	type PhotoResponse struct {
		ThumbnailPhoto string `json:"thumbnail"`
		Photo          string `json:"original"`
	}

	type PetsResponse struct {
		Id            int             `json:"id"`
		PetTypeId     int             `json:"petTypeID"`
		PetTypeName   string          `json:"petType"`
		UserId        uuid.UUID       `json:"userID"`
		Name          string          `json:"petName"`
		BreedId       int             `json:"breedID"`
		BreedName     string          `json:"breed"`
		Photo         []PhotoResponse `json:"photos"`
		BirthDate     time.Time       `json:"birthDate"`
		Male          bool            `json:"male"`
		Gender        string          `json:"gender"`
		Color         string          `json:"color"`
		Care          string          `json:"care"`
		Character     string          `json:"petCharacter"`
		Pedigree      string          `json:"pedigree"`
		Sterilization bool            `json:"sterilization"`
		Vaccinations  bool            `json:"vaccinations"`
	}

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

	if query.Has("userID") {
		userID := query.Get("userID")
		if len(userID) == 0 {
			c.JSON(http.StatusBadRequest, statusResponse{"invalid access token"})
			return
		}
		id, err := uuid.Parse(userID)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		filter.UserId = id
	}

	if query.Has("gender") {
		gender := query.Get("gender")
		if gender == "male" {
			filter.Gender = "male"
		} else if gender == "female" {
			filter.Gender = "female"
		} else {
			newErrorResponse(c, http.StatusInternalServerError, "incorrect gender format")
			return
		}
	}

	if query.Has("petTypeID") {
		PetTypeId, err := strconv.Atoi(query.Get("petTypeID"))
		if err != nil || PetTypeId <= 0 {
			newErrorResponse(c, http.StatusBadRequest, "invalid pet type id param")
			return
		}
		filter.PetTypeId = PetTypeId
	}

	if query.Has("breedID") {
		BreedId, err := strconv.Atoi(query.Get("breedID"))
		if err != nil || BreedId <= 0 {
			newErrorResponse(c, http.StatusBadRequest, "invalid breed id param")
			return
		}
		filter.BreedId = BreedId
	}

	petCardList, err := h.services.PetCard.GetAll(filter)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(petCardList) == 0 {
		newErrorResponse(c, http.StatusOK, "records not found")
		return
	}

	var resp []PetsResponse

	for i := 0; i < len(petCardList); i++ {
		var photos []PhotoResponse
		originalPhoto := strings.Split(petCardList[i].Photo, ", ")
		thumbnailPhoto := strings.Split(petCardList[i].ThumbnailPhoto, ", ")
		for j := 0; j < len(originalPhoto); j++ {
			photos = append(photos, PhotoResponse{ThumbnailPhoto: strings.TrimSpace(thumbnailPhoto[j]),
				Photo: strings.TrimSpace(originalPhoto[j])})
		}
		resp = append(resp,
			PetsResponse{
				Id:            petCardList[i].Id,
				PetTypeId:     petCardList[i].PetTypeId,
				PetTypeName:   petCardList[i].PetTypeName,
				UserId:        petCardList[i].UserId,
				Name:          petCardList[i].Name,
				BreedId:       petCardList[i].BreedId,
				BreedName:     petCardList[i].BreedName,
				Photo:         photos,
				BirthDate:     petCardList[i].BirthDate,
				Male:          petCardList[i].Male,
				Gender:        petCardList[i].Gender,
				Color:         petCardList[i].Color,
				Care:          petCardList[i].Care,
				Character:     petCardList[i].Character,
				Pedigree:      petCardList[i].Pedigree,
				Sterilization: petCardList[i].Sterilization,
				Vaccinations:  petCardList[i].Vaccinations,
			})
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) getMainCardInfo(c *gin.Context) {

	type PetsResponse struct {
		Id             int       `json:"id"`
		PetTypeName    string    `json:"petType"`
		Name           string    `json:"petName"`
		Gender         string    `json:"gender"`
		BreedName      string    `json:"breed"`
		ThumbnailPhoto []string  `json:"thumbnailPhoto"`
		BirthDate      time.Time `json:"birthDate"`
	}

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

	if query.Has("userID") {
		userID := query.Get("userID")
		if len(userID) == 0 {
			c.JSON(http.StatusBadRequest, statusResponse{"invalid access token"})
			return
		}
		id, err := uuid.Parse(userID)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		filter.UserId = id
	}

	if query.Has("gender") {
		gender := query.Get("gender")
		if gender == "male" {
			filter.Gender = "male"
		} else if gender == "female" {
			filter.Gender = "female"
		} else {
			newErrorResponse(c, http.StatusInternalServerError, "incorrect gender format")
			return
		}
	}

	if query.Has("petTypeID") {
		PetTypeId, err := strconv.Atoi(query.Get("petTypeID"))
		if err != nil || PetTypeId <= 0 {
			newErrorResponse(c, http.StatusBadRequest, "invalid pet type id param")
			return
		}
		filter.PetTypeId = PetTypeId
	}

	if query.Has("breedID") {
		BreedId, err := strconv.Atoi(query.Get("breedID"))
		if err != nil || BreedId <= 0 {
			newErrorResponse(c, http.StatusBadRequest, "invalid breed id param")
			return
		}
		filter.BreedId = BreedId
	}

	petCardList, err := h.services.PetCard.GetMain(filter)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(petCardList) == 0 {
		newErrorResponse(c, http.StatusOK, "records not found")
		return
	}

	var resp []PetsResponse

	for i := 0; i < len(petCardList); i++ {
		resp = append(resp,
			PetsResponse{
				Id:             petCardList[i].Id,
				PetTypeName:    petCardList[i].PetTypeName,
				Name:           petCardList[i].Name,
				Gender:         petCardList[i].Gender,
				BreedName:      petCardList[i].BreedName,
				ThumbnailPhoto: strings.Split(petCardList[i].ThumbnailPhoto, ", "),
				BirthDate:      petCardList[i].BirthDate,
			})
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) updateCard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	/*Проверка, что такой pet card id существует*/
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

	/*Проверка, что такой pet type id существует*/
	if input.PetTypeId != nil {
		if _, err := h.services.PetType.GetAll(models.PetTypeFilter{PetTypeId: *input.PetTypeId}); err != nil {
			c.JSON(http.StatusBadRequest, statusResponse{"incorrect pet type id"})
			return
		}
	}

	/*Проверка, что такой breed id существует*/
	if input.BreedId != nil {
		breedId, err := h.services.Breed.GetAll(models.BreedFilter{BreedId: *input.BreedId})
		if err != nil || breedId[0].PetTypeId != *input.PetTypeId {
			c.JSON(http.StatusBadRequest, statusResponse{"incorrect breed id"})
			return
		}
	}

	userID := c.Request.Header.Get("userID")

	if len(userID) == 0 {
		c.JSON(http.StatusBadRequest, statusResponse{"invalid access token"})
		return
	}

	parseUserID, err := uuid.Parse(userID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}

	/*Проверка на то, что id из токена совпадает с id владельца карточки*/
	if petCard[0].UserId != parseUserID {
		newErrorResponse(c, http.StatusBadRequest, "not enough permissions to update")
		return
	}

	input.UserId = &parseUserID

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

	userID := c.Request.Header.Get("userID")
	if len(userID) == 0 {
		c.JSON(http.StatusBadRequest, statusResponse{"invalid access token"})
		return
	}

	parseUserID, err := uuid.Parse(userID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}

	/*Проверка на то, что id из токена совпадает с id владельца карточки*/
	if petCard[0].UserId != parseUserID {
		newErrorResponse(c, http.StatusBadRequest, "not enough permissions to delete")
		return
	}

	err = h.services.PetCard.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
