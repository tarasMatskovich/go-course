package repository

import (
	"library/pkg/model"
)

type BookRepository interface {
	CreateBooks(books []model.Book) (int, error)
	GetBooks() ([]model.Book, error)
}

type bookRepository struct {
	books []model.Book
}

func NewBookRepository() BookRepository {
	return &bookRepository{}
}

func (r *bookRepository) CreateBooks(books []model.Book) (int, error) {
	r.books = books

	return len(r.books), nil
}

func (r *bookRepository) GetBooks() ([]model.Book, error) {
	return r.books, nil
}
