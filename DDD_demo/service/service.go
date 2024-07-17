package service

import (
	"v1/domain"
	"v1/respository"
)

type BookService struct {
	repo respository.BookRespository
}

func NewBookService(repo respository.BookRespository) *BookService {
	return &BookService{
		repo: repo,
	}
}

func (s *BookService) RegisterBook(id, title, author, isbn string) error {
	book := domain.NewBook(id, title, author, isbn)
	return s.repo.Save(book)
}

func (s *BookService) ChangeBookTitle(id, newtitle string) error {
	book, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	book.ChangeTitle(newtitle)
	return s.repo.Save(book)
}
