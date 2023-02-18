package service

import (
	"library/pkg/repository"
	"library/pkg/model"
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