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
			petTypes.GET("/", h.getAllTypes)
		}

		breeds := api.Group("/breeds")
		{
			breeds.GET("/", h.getAllBreeds)
		}

		petCards := api.Group("petCards")
		{
			petCards.POST("/new/", h.createNewCard)
			petCards.GET("/", h.getAllCards)
			petCards.GET("/main/", h.getMainCardInfo)
			petCards.PUT("/update/:id", h.updateCard)
			petCards.DELETE("/delete/:id", h.deleteCard)
		}

		gwConnect := api.Group("endpoint-info")
		{
			gwConnect.GET("/", h.handleInfo)
		}

	}

	h.services.Router = router

	return router
}
