package service_test

import (
	"library/config"
	"library/pkg/mock"
	"library/pkg/model"
	"library/pkg/service"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateBooks(t *testing.T) {
	t.Run("Successfully created list of books", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := mock.NewMockBookRepository(ctrl)
		m.EXPECT().CreateBooks(mock.Books).Return(1, nil)
		service := service.NewBookService(m)

		result, err := service.CreateBooks(mock.BooksList)

		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})

	t.Run("Error occured on creating list of books", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := mock.NewMockBookRepository(ctrl)
		m.EXPECT().CreateBooks(mock.Books).Return(0, mock.ErrorOnCreate)
		service := service.NewBookService(m)

		result, err := service.CreateBooks(mock.BooksList)

		assert.Equal(t, 0, result)
		assert.ErrorIs(t, err, mock.ErrorOnCreate)
	})
}

func TestGetBooks(t *testing.T) {
	t.Run("Successfully got list of books", func(t *testing.T) {
		config := config.Config{
			TimeFormat: "02.01.2006",
		}
		format := config.TimeFormat
		ctrl := gomock.NewController(t)
		m := mock.NewMockBookRepository(ctrl)
		m.EXPECT().GetBooks().Return(mock.Books, nil)
		service := service.NewBookService(m)

		result, err := service.GetBooks()

		assert.Nil(t, err)
		assert.Equal(t, mock.BooksList.Books, result.Books)
		assert.Equal(t, time.Now().Format(format), result.Date.Date.Format(format))

	})

	t.Run("Error on get list of books", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := mock.NewMockBookRepository(ctrl)
		m.EXPECT().GetBooks().Return(nil, mock.ErrorOnGet)
		service := service.NewBookService(m)

		result, err := service.GetBooks()

		assert.Equal(t, model.BooksList{}, result)
		assert.ErrorIs(t, err, mock.ErrorOnGet)
	})
}
