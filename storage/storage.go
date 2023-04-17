package storage

import (
	"context"
	"gorm.io/gorm"
	"repotest/config"
	"repotest/model"
	"repotest/storage/postgre"
)

type IBookRepository interface {
	Get()
	Create()
	Delete()
}

type IUserRepository interface {
	Create(ctx context.Context, user model.User) (uint, error)
	Update(ctx context.Context, user model.User, someField string) error
	Delete(ctx context.Context, ID int) error
	GetByID(ctx context.Context, id uint) (string, error)
	GetAll(ctx context.Context) ([]model.User, error)
	Auth(ctx context.Context, user model.User) error
	GetByUsername(ctx context.Context, username string) (model.User, error)
	CheckIsPhoneExist(ctx context.Context, username string) (bool, error)
	IsVerified(ctx context.Context, username string) (bool, error)
	Verify(ctx context.Context, username string) error
}

type Storage struct {
	pg   *gorm.DB
	Book IBookRepository
	User IUserRepository
}

// Сырые запросы = SQL PGX => PGXPool
// GET users?name="Ruslan"&Surname="Vorontsov"

// ORM = GORM + PoolCon
// Медленнее выполнение присваивания, но быстрая разработка

// CodeGen SQLc
// Под капотом высокопроизводительные драйверы, но много времени уходит на маппинг с основной моделью

func New(ctx context.Context, cfg *config.Config) (*Storage, error) {
	pgDB, err := postgre.Dial(ctx, cfg.Database)
	if err != nil {
		return nil, err
	}

	uRepo := postgre.NewUserRepo(pgDB)

	var storage Storage
	storage.User = uRepo
	return &storage, nil
}
