package marshaller

import (
	"github.com/boon-neko/golang-grpc-backend-example/internal/domain/model"
	admin_apiv1 "github.com/boon-neko/golang-grpc-backend-example/internal/infrastructure/grpc/pb/backend/admin_api/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func StaffToPb(m *model.User) *admin_apiv1.User {
	return &admin_apiv1.User{
		Id:        m.ID,
		Name:      m.Name,
		Role:      uint32(m.Role.Uint()),
		CreatedAt: timestamppb.New(m.CreatedAt),
		UpdatedAt: timestamppb.New(m.UpdatedAt),
	}
}
