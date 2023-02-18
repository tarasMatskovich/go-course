package repository

import (
	"library/pkg/model"
)

type Book interface {
	CreateBooks(users []model.Book) (int, error)
	GetBooks() ([]model.Book, error)
}

type Repository struct {
	Book
}

func NewRepository() *Repository {
	return &Repository{
		Book: NewBookRepository(),
	}
}