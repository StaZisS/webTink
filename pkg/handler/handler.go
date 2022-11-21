package handler

import (
	"github.com/gin-gonic/gin"
	"web/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/log-out", h.logOut)
		auth.POST("/refresh", h.refresh)
	}
	api := router.Group("/api")
	{
		api.GET("/posts", h.getAllPosts)
		api.GET("/posts/:id", h.getPostById)
		api.POST("/posts", h.createPost)
		api.PUT("/posts/:id", h.updatePost)
		api.DELETE("/posts/:id", h.deletePost)
	}
	return router
}