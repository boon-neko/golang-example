package repository

import (
	"github.com/boon-neko/golang-grpc-backend-example/internal/domain/model"
	"golang.org/x/net/context"
)

type UserRepository interface {
	Create(ctx context.Context, m *model.User) error
	Get(ctx context.Context, param *GetUserParam) (*model.User, error)
	GetWithBooks(ctx context.Context, param *GetUserParam) (*model.User, error)
	List(ctx context.Context, param *ListUserParam) (model.Users, error)
	ListWithBooks(ctx context.Context, param *GetUserParam) (*model.User, error)
}

type GetUserParam struct {
	BaseGetOptions
	ID string
}

type ListUserParam struct {
	BaseListOptions
}
