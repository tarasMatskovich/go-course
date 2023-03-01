package mock

import (
	"errors"
	"library/pkg/model"
	"time"
)

var (
	Books []model.Book = []model.Book{
		{
			Name:   "Book name",
			Author: "Book Author",
			Year:   1998,
		},
	}
	BooksList model.BooksList = model.BooksList{
		Books: Books,
		Date: model.BookTime{Date: time.Now()},
	}
	ErrorOnCreate = errors.New("error on creating list of books")
	ErrorOnGet    = errors.New("error on get list of books")
)
