package handler

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

func (h *Handler) InitRoutes() *gin.Engine {
	h.routes = gin.New()
	h.routes.Use(cors.Default())

	auth := h.routes.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-out", h.signOut)
	}

	return h.routes
}
