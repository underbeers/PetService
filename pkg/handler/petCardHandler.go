package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/underbeers/PetService/pkg/models"
	"io"
	"net/http"
	"strconv"
	"strings"
	_ "strings"
	"time"
)

const (
	userIDAuth = "UserID"
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

	input.Photo = "{https://res.cloudinary.com/dojhrhddc/image/upload/v1684863033/pet_placeholder_rcf7iv.jpg}"
	input.ThumbnailPhoto = "{https://res.cloudinary.com/dojhrhddc/image/upload/v1684864098/pet_placeholder_rcf7iv_i1lzt9.jpg}"

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

	var resp []PetsResponse

	if len(petCardList) == 0 {
		resp = make([]PetsResponse, 0)
	}

	for i := 0; i < len(petCardList); i++ {
		var photos []PhotoResponse

		if len(petCardList[i].Photo) > 5 && len(petCardList[i].ThumbnailPhoto) > 5 {
			originalPhoto := strings.Split(petCardList[i].Photo[1:len(petCardList[i].Photo)-1], ",")
			thumbnailPhoto := strings.Split(petCardList[i].ThumbnailPhoto[1:len(petCardList[i].ThumbnailPhoto)-1], ",")
			//Если фото добавлено, то не учитываем дефолтное фото
			if len(originalPhoto) > 1 {
				originalPhoto = originalPhoto[1:]
			}
			if len(thumbnailPhoto) > 1 {
				thumbnailPhoto = thumbnailPhoto[1:]
			}
			for j := 0; j < len(originalPhoto) || j < len(thumbnailPhoto); j++ {
				photos = append(photos, PhotoResponse{})
				if j < len(thumbnailPhoto) {
					photos[j].ThumbnailPhoto = strings.TrimSpace(thumbnailPhoto[j])
				}
				if j < len(originalPhoto) {
					photos[j].Photo = strings.TrimSpace(originalPhoto[j])
				}
			}
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
		ThumbnailPhoto string    `json:"photo"`
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

	var resp []PetsResponse

	if len(petCardList) == 0 {
		resp = make([]PetsResponse, 0)
	}

	for i := 0; i < len(petCardList); i++ {
		thumbnailPhoto := strings.Split(petCardList[i].ThumbnailPhoto[1:len(petCardList[i].ThumbnailPhoto)-1], ",")
		if len(thumbnailPhoto) > 1 {
			thumbnailPhoto = thumbnailPhoto[1:]
		}
		resp = append(resp,
			PetsResponse{
				Id:             petCardList[i].Id,
				PetTypeName:    petCardList[i].PetTypeName,
				Name:           petCardList[i].Name,
				Gender:         petCardList[i].Gender,
				BreedName:      petCardList[i].BreedName,
				ThumbnailPhoto: thumbnailPhoto[0],
				BirthDate:      petCardList[i].BirthDate,
			})
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) updateCard(c *gin.Context) {
	var id int
	query := c.Request.URL.Query()
	if query.Has("id") {
		cardID, err := strconv.Atoi(query.Get("id"))
		id = cardID
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "invalid id param")
			return
		}
	} else {
		newErrorResponse(c, http.StatusBadRequest, "id not provided")
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

func (h *Handler) setImage(c *gin.Context) {
	type Data struct {
		Original  string `json:"original"`
		Thumbnail string `json:"thumbnail"`
	}

	type Request struct {
		StatusCode int  `json:"statusCode"`
		Data       Data `json:"data"`
		PetCardID  int  `json:"petCardID"`
	}

	req := &Request{}
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	if err := json.Unmarshal(body, req); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	err = h.services.PetCard.SetImage(req.PetCardID, req.Data.Thumbnail, req.Data.Original)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) transferPet(c *gin.Context) {

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

	var input models.TransferPet
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	/*Проверка, что такой pet card id существует*/
	petCard, err := h.services.PetCard.GetAll(models.PetCardFilter{PetCardId: input.Id})
	if len(petCard) != 1 || err != nil {
		c.JSON(http.StatusBadRequest, statusResponse{"incorrect pet card id"})
		return
	}

	/*Проверка на то, что id из токена совпадает с id владельца карточки*/
	if petCard[0].UserId != parseUserID {
		newErrorResponse(c, http.StatusBadRequest, "not enough permissions to transfer pet card")
		return
	}

	if petCard[0].UserId == input.NewOwnerID {
		newErrorResponse(c, http.StatusBadRequest, "ID of new and current owners are the same")
		return
	}

	if err := h.services.PetCard.TransferPet(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteCard(c *gin.Context) {
	var id int
	query := c.Request.URL.Query()
	if query.Has("id") {
		cardID, err := strconv.Atoi(query.Get("id"))
		id = cardID
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "invalid id param")
			return
		}
	} else {
		newErrorResponse(c, http.StatusBadRequest, "id not provided")
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
