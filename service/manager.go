package service

import (
	"errors"
	"repotest/storage"
)

type Manager struct {
	Book IBook
	User IUserService
}

func NewManager(storage *storage.Storage) (*Manager, error) {
	uSrv := NewUserService(storage)

	bSrv := NewBookService(storage, *uSrv)
	if storage == nil {
		return nil, errors.New("no storage provided")
	}

	return &Manager{
		Book: bSrv,
		//User: uSrv,
	}, nil
}
