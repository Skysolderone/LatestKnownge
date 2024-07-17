package respository

import "v1/domain"

type BookRespository interface {
	Save(book *domain.Book) error
	FindByID(id string) (*domain.Book, error)
	Remove(id string) error
}
