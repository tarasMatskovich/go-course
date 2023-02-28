package model

import (
	"encoding/json"
	"strconv"
)

type Book struct {
	Name   string `json:"name" binding:"required"`
	Author string `json:"author" binding:"required"`
	Year   int    `json:"year" binding:"required"`
}

type BooksList struct {
	Books []Book   `json:"books" binding:"required"`
	Date  BookTime `json:"date" binding:"required"`
}

type SortedBooks []Book

func (b SortedBooks) Len() int {
	return len(b)
}
func (b SortedBooks) Less(i, j int) bool {
	return b[i].Year < b[j].Year
}
func (b SortedBooks) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b *Book) MarshalJSON() ([]byte, error) {
	data := struct {
		Name   string `json:"name"`
		Author string `json:"author"`
		Year   string `json:"year"`
	}{
		Name:   b.Name,
		Author: b.Author,
		Year:   strconv.Itoa(b.Year),
	}

	return json.Marshal(&data)
}

func (b *Book) UnmarshalJSON(data []byte) error {
	book := struct {
		Name   string
		Author string
		Year   string
	}{}

	if err := json.Unmarshal(data, &book); err != nil {
		return err
	}

	var err error
	b.Name = book.Name
	b.Author = book.Author
	b.Year, err = strconv.Atoi(book.Year)

	return err
}
