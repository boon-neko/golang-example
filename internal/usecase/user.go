package usecase

import (
	"context"
	"github.com/boon-neko/golang-grpc-backend-example/internal/domain/model"
	"github.com/boon-neko/golang-grpc-backend-example/internal/usecase/input"
)

type UserUsecase interface {
	CreateUser(ctx context.Context)
	GetUser(
		ctx context.Context,
		param *input.GetUserRequest,
	) (*model.User, error)
}
