package handler

import (
	"library/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateBooks(c *gin.Context) {
	var input []model.Book

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.services.Book.CreateBooks(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "success",
	})
}

func (h *Handler) GetBooks(c *gin.Context) {
	books, err := h.services.Book.GetBooks()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, books)
}
