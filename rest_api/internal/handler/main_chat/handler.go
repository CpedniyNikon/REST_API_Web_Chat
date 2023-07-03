package main_chat

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	routes *gin.Engine
}

func NewHandler() (h *Handler) {
	return &Handler{}
}

func (h *Handler) InitMainChatRoutes() *gin.Engine {
	h.routes = gin.New()
	h.routes.Use(cors.Default())
	auth := h.routes.Group("/chat")
	{
		auth.POST("/sign-out", h.write)
	}
	return h.routes
}
