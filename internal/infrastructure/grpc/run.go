package grpc

import (
	"github.com/boon-neko/golang-grpc-backend-example/internal/infrastructure/di"
	"github.com/boon-neko/golang-grpc-backend-example/internal/infrastructure/grpc/internal/handler/admin"
	admin_apiv1 "github.com/boon-neko/golang-grpc-backend-example/internal/infrastructure/grpc/pb/backend/admin_api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewServer(box *di.DIBox) *grpc.Server {
	server := grpc.NewServer(
	// grpc.KeepaliveEnforcementPolicy(
	//	keepalive.EnforcementPolicy{
	//		MinTime:             10 * time.Second,
	//		PermitWithoutStream: true,
	//	},
	//),
	//grpc.KeepaliveParams(
	//	keepalive.ServerParameters{
	//		MaxConnectionIdle:     0,
	//		MaxConnectionAge:      10 * time.Minute,
	//		MaxConnectionAgeGrace: 0,
	//		Time:                  20 * time.Second,
	//		Timeout:               10 * time.Second,
	//	},
	)

	adminHandler := admin.NewHandler(box.UserUsecase)

	admin_apiv1.RegisterAdminV1ServiceServer(server, adminHandler)

	reflection.Register(server)
	return server
}
