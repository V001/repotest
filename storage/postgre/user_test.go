package postgre

import (
	"context"
	"log"
	"repotest/config"
	"repotest/model"
	"testing"
)

func TestUserRepo_Create(t *testing.T) {
	type args struct {
		ctx  context.Context
		user model.User
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr bool
	}{
		{
			name:    "Success",
			args:    args{context.Background(), model.User{ID: 3}},
			want:    3,
			wantErr: true,
		},
	}
	repo := CreateUserRepo(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.Create(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}

			name, err := repo.GetByID(context.Background(), got)
			if err != nil {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
			log.Print(name)
		})
	}
}

func CreateUserRepo(t *testing.T) *UserRepo {
	cfg, _ := config.New()
	db, _ := Dial(context.Background(), cfg.Database)
	return NewUserRepo(db)
}
