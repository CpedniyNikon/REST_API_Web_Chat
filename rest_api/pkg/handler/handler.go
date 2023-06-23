package handler

import (
	"github.com/gin-gonic/gin"
	_ "net/http"
)

type Handler struct {
	routes *gin.Engine
}

func NewHandler() (h *Handler) {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	h.routes = gin.New()

	auth := h.routes.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	return h.routes
}
