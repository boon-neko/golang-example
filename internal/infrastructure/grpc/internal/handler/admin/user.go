package admin

import (
	"context"
	"github.com/boon-neko/golang-grpc-backend-example/internal/infrastructure/grpc/internal/handler/admin/marshaller"
	admin_apiv1 "github.com/boon-neko/golang-grpc-backend-example/internal/infrastructure/grpc/pb/backend/admin_api/v1"
	"github.com/boon-neko/golang-grpc-backend-example/internal/usecase/input"
)

func (h *Handler) GetUser(ctx context.Context, request *admin_apiv1.GetUserRequest) (*admin_apiv1.GetUserResponse, error) {
	got, err := h.userInteractor.GetUser(ctx, &input.GetUserRequest{ID: request.GetId()})
	if err != nil {
		return nil, err
	}

	return &admin_apiv1.GetUserResponse{
		User: marshaller.StaffToPb(got),
	}, nil
}

func (h *Handler) CreateUser(ctx context.Context, request *admin_apiv1.CreateUserRequest) (*admin_apiv1.CreateUserResponse, error) {
	//TODO implement me
	panic("implement me")
}
