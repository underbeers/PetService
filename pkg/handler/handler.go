package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/underbeers/PetService/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		petTypes := api.Group("/petTypes")
		{
			petTypes.GET("/", h.getAllTypes)
		}

		breeds := api.Group("/breeds")
		{
			breeds.GET("/", h.getAllBreeds)
		}

		petCards := api.Group("petCards")
		{
			petCards.POST("/", h.createNewCard)
			petCards.GET("/", h.getAllCards)
			petCards.GET("/main/", h.getMainCardInfo)
			petCards.PUT("/:id", h.updateCard)
			petCards.DELETE("/:id", h.deleteCard)
		}

	}

	return router
}
