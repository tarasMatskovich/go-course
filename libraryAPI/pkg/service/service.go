package service

import (
	"library/pkg/model"
	"library/pkg/repository"
)

type Book interface {
	CreateBooks(users []model.Book) (int, error)
	GetBooks() ([]model.Book, error)
}

type Service struct {
	Book
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Book: NewBookService(repos.Book),
	}
}
