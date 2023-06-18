package handler

import (
	"github.com/gin-gonic/gin"
	_ "net/http"
	"rest_api/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) (h *Handler) {
	return &Handler{services}
}

func (h *Handler) InitRoutes() *gin.Engine  {
	routes := gin.New()
	auth:= routes.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api:= routes.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.createList)
			lists.DELETE("/:id", h.deleteList)

			items := api.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/",h.getAllItems)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.createItem)
				items.DELETE("/:item_id", h.deleteItem)
			}

		}

	}

	return routes
}
