package service

import (
	"library/pkg/model"
	"library/pkg/repository"
)

type Book interface {
	CreateBooks(books model.BooksList) (int, error)
	GetBooks() (model.BooksList, error)
}

type Service struct {
	Book
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Book: NewBookService(repos.BookRepository),
	}
}
