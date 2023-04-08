package service

import (
	"context"
	"repotest/storage"
)

type IBook interface {
	Get(ctx context.Context)
	Create()
}

type BookService struct {
	Repo *storage.Storage
	User UserService
}

func NewBookService(repo *storage.Storage, user UserService) *BookService {
	return &BookService{Repo: repo, User: user}
}

func (s *BookService) Get(ctx context.Context) {

	s.Repo.Book.Get()
}

func (s *BookService) Create() {
	//uuid := uuid2.NewGen()
	s.Repo.Book.Create()
}
