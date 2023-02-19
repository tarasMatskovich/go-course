package repository

type Repository struct {
	BookRepository
}

func NewRepository() *Repository {
	return &Repository{
		BookRepository: NewBookRepository(),
	}
}
