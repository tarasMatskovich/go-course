package service_test

import (
	"library/pkg/mock"
	"library/pkg/service"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateBooks(t *testing.T) {
	t.Run("Successfully created list of books", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := mock.NewMockBookRepository(ctrl)
		m.EXPECT().CreateBooks(mock.Books).Return(1, nil)
		service := service.NewBookService(m)

		result, err := service.CreateBooks(mock.Books)

		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})

	t.Run("Error occured on creating list of books", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := mock.NewMockBookRepository(ctrl)
		m.EXPECT().CreateBooks(mock.Books).Return(0, mock.ErrorOnCreate)
		service := service.NewBookService(m)

		result, err := service.CreateBooks(mock.Books)

		assert.Equal(t, 0, result)
		assert.ErrorIs(t, err, mock.ErrorOnCreate)
	})
}

func TestGetBooks(t *testing.T) {
	t.Run("Successfully got list of books", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := mock.NewMockBookRepository(ctrl)
		m.EXPECT().GetBooks().Return(mock.Books, nil)
		service := service.NewBookService(m)

		result, err := service.GetBooks()

		assert.Nil(t, err)
		assert.Equal(t, mock.Books, result)
	})

	t.Run("Error on get list of books", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := mock.NewMockBookRepository(ctrl)
		m.EXPECT().GetBooks().Return(nil, mock.ErrorOnGet)
		service := service.NewBookService(m)

		result, err := service.GetBooks()

		assert.Nil(t, result)
		assert.ErrorIs(t, err, mock.ErrorOnGet)
	})
}
