package admin

import (
	admin_apiv1 "github.com/boon-neko/golang-grpc-backend-example/internal/infrastructure/grpc/pb/backend/admin_api/v1"
	"github.com/boon-neko/golang-grpc-backend-example/internal/usecase"
)

type Handler struct {
	userInteractor usecase.UserUsecase
}

func NewHandler(userInteractor usecase.UserUsecase) admin_apiv1.AdminV1ServiceServer {
	return &Handler{userInteractor: userInteractor}
}
