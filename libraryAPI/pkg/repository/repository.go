package repository

import (
	"library/pkg/storage"
)

type Repository struct {
	BookRepository
}

func NewRepository(bookStorage *storage.Storage) *Repository {
	return &Repository{
		BookRepository: NewCSVBookRepository(bookStorage),
	}
}
