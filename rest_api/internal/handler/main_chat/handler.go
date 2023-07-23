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
	chat := h.routes.Group("/chat")
	{
		chat.POST("/write", h.write)
		chat.POST("/get_messages", h.getMessages)
	}
	return h.routes
}
