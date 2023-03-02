package service

import (
	"library/pkg/model"
	"library/pkg/repository"
	"sort"
	"time"
)

type BookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) CreateBooks(books model.BooksList) (int, error) {
	return s.repo.CreateBooks(books.Books)
}

func (s *BookService) GetBooks() (model.BooksList, error) {
	books, err := s.repo.GetBooks()
	if err != nil {
		return model.BooksList{}, err
	}

	sort.Sort(model.SortedBooks(books))

	date := &model.BookTime{
		Date: time.Now(),
	}

	return model.BooksList{
		Books: books,
		Date:  *date,
	}, nil
}
