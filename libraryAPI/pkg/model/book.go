package model

type Book struct {
	Name   string `json:"name" binding:"required"`
	Author string `json:"author" binding:"required"`
	Year   int    `json:"year,string" binding:"required"`
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
