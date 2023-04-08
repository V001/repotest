package service

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"repotest/model"
	"repotest/storage"
	"time"
)

type IUserService interface {
	Get()
	Create(ctx context.Context, req model.UserCreateReq) (model.CreateResp, error)
	CheckPassword(encPass, providedPassword string) error
	HashPassword(password string) (string, error)
	RefreshToken() (string, error)
	Auth(ctx context.Context, user model.AuthUser) error
}

type UserService struct {
	repo *storage.Storage
}

func NewUserService(repo *storage.Storage) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Get() {
	//TODO implement me
	panic("implement me")
}

func (s *UserService) Create(ctx context.Context, user model.User) (uint, error) {
	var err error
	user.Password, err = s.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	return s.repo.User.Create(ctx, user)
}

func (s *UserService) CheckPassword(encPass, providedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(encPass), []byte(providedPassword))
}

func (s *UserService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (s *UserService) Auth(ctx context.Context, user model.AuthUser) error {
	userFromDB, userErr := s.repo.User.GetByUsername(ctx, user.Username)
	if userErr != nil {
		return userErr
	}
	checkErr := s.CheckPassword(userFromDB.Password, user.Password)
	if checkErr != nil {
		return checkErr
	}
	return nil
}

func (s *UserService) RefreshToken() (string, error) {
	b := make([]byte, 32)

	str := rand.NewSource(time.Now().Unix())
	r := rand.New(str)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

func (s *UserService) Update(ctx context.Context, user model.User) error {
	if user.ID == 0 && len(user.Username) == 0 {
		return fmt.Errorf("empty ID")
	}
	if len(user.Password) != 0 {
		user.Password, _ = s.HashPassword(user.Password)
	}
	return s.repo.User.Update(ctx, user)
}
