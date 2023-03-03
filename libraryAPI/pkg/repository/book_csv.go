package repository

import (
	"errors"
	"library/pkg/model"
	"library/pkg/storage"

	"github.com/gocarina/gocsv"
)

type bookCSVRepository struct {
	storage *storage.Storage
}

func NewCSVBookRepository(storage *storage.Storage) BookRepository {
	return &bookCSVRepository{
		storage: storage,
	}
}

func (r *bookCSVRepository) CreateBooks(books []model.Book) (int, error) {
	if err := r.storage.ClearPointer(); err != nil {
		return 0, err
	}

	booksFromFile := []model.Book{}
	if err := gocsv.UnmarshalFile(r.storage.File, &booksFromFile); err != nil {
		if !errors.Is(err, gocsv.ErrEmptyCSVFile) {
			return 0, err
		}
	}

	if err := r.storage.ClearPointer(); err != nil {
		return 0, err
	}

	books = append(books, booksFromFile...)

	err := gocsv.MarshalFile(&books, r.storage.File)
	if err != nil {
		return 0, err
	}

	return len(books), nil
}

func (r *bookCSVRepository) GetBooks() ([]model.Book, error) {
	if err := r.storage.ClearPointer(); err != nil {
		return nil, err
	}

	books := []model.Book{}

	if err := gocsv.UnmarshalFile(r.storage.File, &books); err != nil {
		if errors.Is(err, gocsv.ErrEmptyCSVFile) {
			return nil, nil
		}

		return nil, err
	}

	return books, nil
}
