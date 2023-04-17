package postgre

import (
	"context"
	"gorm.io/gorm"
	"repotest/model"
)

type UserRepo struct {
	DB *gorm.DB
}

func (r *UserRepo) CheckIsPhoneExist(ctx context.Context, username string) (bool, error) {
	var cnt int64
	err := r.DB.WithContext(ctx).Table("users").Where("username = ?", username).Count(&cnt).Error
	return cnt > 0, err
}

func (r *UserRepo) Verify(ctx context.Context, username string) error {
	return r.DB.WithContext(ctx).Table("users").
		Where("username = ?", username).
		UpdateColumn("is_verify", true).Error
}

func (r *UserRepo) IsVerified(ctx context.Context, username string) (bool, error) {
	var resp bool
	err := r.DB.WithContext(ctx).
		Table("users").
		Select("is_verify").
		Where("username = ?", username).
		Find(&resp).Error
	return true, err

}

func NewUserRepo(DB *gorm.DB) *UserRepo {
	return &UserRepo{DB: DB}
}

// 5sec
// 150ms

func (r *UserRepo) Create(ctx context.Context, user model.User) (uint, error) {
	result := r.DB.WithContext(ctx).Omit("deleted_at").Create(&user)
	return user.ID, result.Error
}

func (r *UserRepo) GetByID(ctx context.Context, id uint) (string, error) {
	var resp string
	result := r.DB.Table("users").WithContext(ctx).Where("id = ?", id).Select("name").Find(&resp)
	return resp, result.Error
}

func (r *UserRepo) Update(ctx context.Context, user model.User, someField string) error {
	return r.DB.Where("username = ?", user.Name).Updates(&user).Error
}

func (r *UserRepo) Delete(ctx context.Context, ID int) error {
	return r.DB.WithContext(ctx).Delete(&model.User{}, ID).Error
}

func (r *UserRepo) GetAll(ctx context.Context) ([]model.User, error) {
	var resp []model.User
	err := r.DB.WithContext(ctx).Find(&resp)
	return resp, err.Error
}

func (r *UserRepo) Auth(ctx context.Context, user model.User) error {
	panic("implement me")
	return nil
}

func (r *UserRepo) GetByUsername(ctx context.Context, username string) (model.User, error) {
	var res model.User
	err := r.DB.WithContext(ctx).Where("username = ?", username).First(&res).Error
	if err != nil {
		return model.User{}, err
	}
	return res, nil
}
