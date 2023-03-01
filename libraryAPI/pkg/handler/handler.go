package handler

import (
	"library/pkg/middleware"
	"library/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.JSONMiddleware())

	router.POST("/books", h.CreateBooks)
	router.GET("/books", h.GetBooks)

	return router
}
