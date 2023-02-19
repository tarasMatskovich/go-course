package service

import (
	"library/pkg/model"
	"library/pkg/repository"
	"sort"
)

type BookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) CreateBooks(books []model.Book) (int, error) {
	sort.Sort(model.SortedBooks(books))
	return s.repo.CreateBooks(books)
}

func (s *BookService) GetBooks() ([]model.Book, error) {
	return s.repo.GetBooks()
}
