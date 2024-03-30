package repository

import (
	"github.com/boon-neko/golang-grpc-backend-example/internal/domain/model"
	"github.com/boon-neko/golang-grpc-backend-example/internal/domain/repository"
	"golang.org/x/net/context"
	"time"
)

type user struct{}

func NewUser() repository.UserRepository {
	return &user{}
}

func (u user) Create(ctx context.Context, m *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (u user) Get(ctx context.Context, param *repository.GetUserParam) (*model.User, error) {
	return &model.User{
		ID:        param.ID,
		Name:      "test",
		Role:      0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}, nil
}

func (u user) GetWithBooks(ctx context.Context, param *repository.GetUserParam) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u user) List(ctx context.Context, param *repository.ListUserParam) (model.Users, error) {
	//TODO implement me
	panic("implement me")
}

func (u user) ListWithBooks(ctx context.Context, param *repository.GetUserParam) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}
