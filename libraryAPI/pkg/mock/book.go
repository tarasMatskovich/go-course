package mock

import (
	"errors"
	"library/pkg/model"
)

var (
	Books []model.Book = []model.Book{
		{
			Name:   "Book name",
			Author: "Book Author",
			Year:   1998,
		},
	}
	ErrorOnCreate = errors.New("error on creating list of books")
	ErrorOnGet    = errors.New("error on get list of books")
)
