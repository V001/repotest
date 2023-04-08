package redis

import "repotest/model"

type UserRepository struct {
}

func (r *UserRepository) Get() {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepository) Create(m model.UserCreateReq) (model.CreateResp, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}
