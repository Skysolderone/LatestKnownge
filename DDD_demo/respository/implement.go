package respository

import (
	"errors"

	"v1/domain"
)

type InMemoryBookRespository struct {
	books map[string]*domain.Book
}

func NewInMemoryBookRepository() *InMemoryBookRespository {
	return &InMemoryBookRespository{
		books: make(map[string]*domain.Book),
	}
}

func (r *InMemoryBookRespository) FindByID(id string) (*domain.Book, error) {
	if book, exists := r.books[id]; exists {
		return book, nil
	}
	return nil, errors.New("book not found")
}

func (r *InMemoryBookRespository) Remove(id string) error {
	delete(r.books, id)
	return nil
}
