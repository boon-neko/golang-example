package usecase

import (
	"context"
	"fmt"
	"github.com/boon-neko/golang-grpc-backend-example/internal/domain/model"
	"github.com/boon-neko/golang-grpc-backend-example/internal/domain/repository"
	"github.com/boon-neko/golang-grpc-backend-example/internal/usecase/input"
)

type UserInteractor struct {
	repo repository.UserRepository
}

func NewUserInteractor(repo repository.UserRepository) UserUsecase {
	return &UserInteractor{repo: repo}
}

func (u *UserInteractor) CreateUser(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (u *UserInteractor) GetUser(ctx context.Context, param *input.GetUserRequest) (*model.User, error) {
	fmt.Printf("usecase. request param %#v", param)

	user, err := u.repo.Get(ctx, &repository.GetUserParam{
		BaseGetOptions: repository.BaseGetOptions{},
		ID:             param.ID,
	})

	return user, err
}
