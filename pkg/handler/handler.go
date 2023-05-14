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

	api := router.Group("/api/v1")
	{
		petTypes := api.Group("/petTypes")
		{
			petTypes.GET("", h.getAllTypes).OPTIONS("", h.getAllTypes)
		}

		breeds := api.Group("/breeds")
		{
			breeds.GET("", h.getAllBreeds).OPTIONS("", h.getAllBreeds)
		}

		petCards := api.Group("petCards")
		{
			petCards.POST("/new", h.createNewCard).OPTIONS("/new", h.createNewCard)
			petCards.GET("", h.getAllCards).OPTIONS("", h.getAllCards)
			petCards.GET("/main", h.getMainCardInfo).OPTIONS("/main", h.getMainCardInfo)
			petCards.PUT("/update", h.updateCard).OPTIONS("/update", h.updateCard)
			petCards.DELETE("/delete", h.deleteCard).OPTIONS("/delete", h.deleteCard)

		}

		gwConnect := api.Group("endpoint-info")
		{
			gwConnect.GET("/", h.handleInfo).OPTIONS("/", h.handleInfo)
		}

	}

	h.services.Router = router

	return router
}
