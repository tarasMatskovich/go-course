package handler_test

import (
	"library/pkg/handler"
	"library/pkg/mock"
	"library/pkg/repository"
	"library/pkg/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateBooksHandler(t *testing.T) {
	t.Run("Successfully created list of books", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		ctrl := gomock.NewController(t)
		m := mock.NewMockBookRepository(ctrl)
		m.EXPECT().CreateBooks(mock.Books).Return(1, nil)
		handler := handler.NewHandler(service.NewService(&repository.Repository{
			BookRepository: m,
		}))
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		c.Request = &http.Request{
			Header: make(http.Header),
		}
		mock.MockCreateBooksPostJson(c, mock.Books)

		handler.CreateBooks(c)

		assert.Equal(t, http.StatusCreated, rr.Result().StatusCode)
	})

	t.Run("Error on creating list of books", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		ctrl := gomock.NewController(t)
		m := mock.NewMockBookRepository(ctrl)
		m.EXPECT().CreateBooks(mock.Books).Return(0, mock.ErrorOnCreate)
		handler := handler.NewHandler(service.NewService(&repository.Repository{
			BookRepository: m,
		}))
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		c.Request = &http.Request{
			Header: make(http.Header),
		}
		mock.MockCreateBooksPostJson(c, mock.Books)

		handler.CreateBooks(c)

		assert.Equal(t, http.StatusInternalServerError, rr.Result().StatusCode)
	})

	t.Run("Bad request on creating list of books", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		ctrl := gomock.NewController(t)
		m := mock.NewMockBookRepository(ctrl)
		handler := handler.NewHandler(service.NewService(&repository.Repository{
			BookRepository: m,
		}))
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		c.Request = &http.Request{
			Header: make(http.Header),
		}
		mock.MockCreateBooksPostJson(c, gin.H{"data": "value"})

		handler.CreateBooks(c)

		assert.Equal(t, http.StatusBadRequest, rr.Result().StatusCode)
	})
}

func TestGetBooksHandler(t *testing.T) {
	t.Run("Successfully got list of books", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		ctrl := gomock.NewController(t)
		m := mock.NewMockBookRepository(ctrl)
		m.EXPECT().GetBooks().Return(mock.Books, nil)
		handler := handler.NewHandler(service.NewService(&repository.Repository{
			BookRepository: m,
		}))
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		c.Request = &http.Request{
			Header: make(http.Header),
		}

		handler.GetBooks(c)

		assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	})

	t.Run("Error on get list of books", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		ctrl := gomock.NewController(t)
		m := mock.NewMockBookRepository(ctrl)
		m.EXPECT().GetBooks().Return(nil, mock.ErrorOnGet)
		handler := handler.NewHandler(service.NewService(&repository.Repository{
			BookRepository: m,
		}))
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		c.Request = &http.Request{
			Header: make(http.Header),
		}

		handler.GetBooks(c)

		assert.Equal(t, http.StatusInternalServerError, rr.Result().StatusCode)
	})
}
