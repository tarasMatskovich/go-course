package repository

import (
	"library/pkg/model"
)

type BookRepository struct {
	books []model.Book
}

func NewBookRepository() *BookRepository {
	return &BookRepository{}
}

func (r *BookRepository) CreateBooks(books []model.Book) (int, error) {
	r.books = books
	
	return len(r.books), nil
}

func (r *BookRepository) GetBooks() ([]model.Book, error) {
	return r.books, nil
}