package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"repotest/model"
	"repotest/storage"
	mock_storage "repotest/storage/mocks"
	"testing"
)

// E2У
// integration
// Unit

func TestUserService_Create(t *testing.T) {
	type args struct {
		ctx  context.Context
		user model.User
	}
	tests := []struct {
		prepare func(f *mock_storage.MockIUserRepository)
		name    string
		args    args
		want    uint
		wantErr bool
	}{{
		name:    "success",
		args:    args{context.Background(), model.User{Email: "1"}},
		want:    1,
		wantErr: false,
		prepare: func(f *mock_storage.MockIUserRepository) {
			f.EXPECT().Create(gomock.Any(), model.User{Email: "1"}).
				Return(uint(1), nil)
		},
	},
	}
	for _, tt := range tests {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		userRepo := mock_storage.NewMockIUserRepository(ctrl)
		tt.prepare(userRepo)
		s := NewUserService(&storage.Storage{User: userRepo})
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Create(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFoo(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{1, 2},
			want: 3,
		},
		{
			name: "fail",
			args: args{1, 5},
			want: 6,
		},
		{
			name: "success",
			args: args{1, 2},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Foo(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Foo() = %v, want %v", got, tt.want)
			}
		})
	}
}
