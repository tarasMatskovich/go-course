package handler_test

import (
	"bytes"
	"encoding/json"
	"library/config"
	"library/pkg/handler"
	"library/pkg/mock"
	"library/pkg/repository"
	"library/pkg/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateBooksHandler(t *testing.T) {
	configPath := "./../../configs/config.env"
	config.New(configPath)
	
	jsonResponseBody := []byte(`
	{
		"books": [
			{
				"name": "Book name",
				"author": "Book Author",
				"year": "1998"
			}
		],
		"date": "28.02.2023"
	}
	`)
	failedResponseBody := []byte(`{"data":"value"}`)
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
		mock.MockCreateBooksPostJson(c, jsonResponseBody)

		handler.CreateBooks(c)

		assert.Equal(t, http.StatusCreated, rr.Result().StatusCode)
		expected := `{"status":"success"}`
		assert.Equal(t, expected, rr.Body.String())
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
		mock.MockCreateBooksPostJson(c, jsonResponseBody)

		handler.CreateBooks(c)

		assert.Equal(t, http.StatusInternalServerError, rr.Result().StatusCode)
		expected := `{"message":"error on creating list of books"}`
		assert.Equal(t, expected, rr.Body.String())
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
		mock.MockCreateBooksPostJson(c, failedResponseBody)

		handler.CreateBooks(c)

		assert.Equal(t, http.StatusBadRequest, rr.Result().StatusCode)
		expected := `{"message":"Key: 'BooksList.Books' Error:Field validation for 'Books' failed on the 'required' tag"}`
		assert.Equal(t, expected, rr.Body.String())
	})
}

func TestGetBooksHandler(t *testing.T) {
	configPath := "./../../configs/config.env"
	config.New(configPath)

	body := `
	{
		"books": [
			{
				"name": "Book name",
				"author": "Book Author",
				"year": "1998"
			}
		],
		"date": "{now}"
	}
	`
	now := time.Now().Format("02.01.2006")
	body = strings.Replace(body, "{now}", now, -1)
	responseBody := []byte(body)
	jsonResponseBody := &bytes.Buffer{}
	err := json.Compact(jsonResponseBody, responseBody)
	assert.Nil(t, err)

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
		assert.Equal(t, jsonResponseBody.Bytes(), rr.Body.Bytes())
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
		expected := `{"message":"error on get list of books"}`
		assert.Equal(t, expected, rr.Body.String())
	})
}
