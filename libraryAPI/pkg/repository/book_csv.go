package repository

import (
	"errors"
	"library/pkg/model"
	"os"

	"github.com/gocarina/gocsv"
)

type bookCSVRepository struct {
	file string
}

func NewCSVBookRepository(file string) BookRepository {
	return &bookCSVRepository{
		file: file,
	}
}

func (r *bookCSVRepository) CreateBooks(books []model.Book) (int, error) {
	booksFile, err := os.OpenFile(r.file, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}

	defer booksFile.Close()

	booksFromFile := []model.Book{}
	if err := gocsv.UnmarshalFile(booksFile, &booksFromFile); err != nil {
		if !errors.Is(err, gocsv.ErrEmptyCSVFile) {
			return 0, err
		}
	}

	if _, err := booksFile.Seek(0, 0); err != nil {
		return 0, err
	}

	books = append(books, booksFromFile...)

	err = gocsv.MarshalFile(&books, booksFile)
	if err != nil {
		return 0, err
	}

	return len(books), nil
}

func (r *bookCSVRepository) GetBooks() ([]model.Book, error) {
	booksFile, err := os.OpenFile(r.file, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}

	defer booksFile.Close()

	books := []model.Book{}

	if err := gocsv.UnmarshalFile(booksFile, &books); err != nil {
		if errors.Is(err, gocsv.ErrEmptyCSVFile) {
			return []model.Book{}, nil
		}

		return nil, err
	}

	return books, nil
}
