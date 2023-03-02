package repository

import "library/config"

type Repository struct {
	BookRepository
}

func NewRepository(c config.Config) *Repository {
	return &Repository{
		BookRepository: NewCSVBookRepository(c.RepoFilePath),
	}
}
