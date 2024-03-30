package input

import "github.com/boon-neko/golang-grpc-backend-example/internal/pkg/validation"

type GetUserRequest struct {
	ID string `validate:"required"`
}

func NewGetUserRequest(ID string) *GetUserRequest {
	return &GetUserRequest{ID: ID}
}

func (r *GetUserRequest) Validate() error {
	if err := validation.Validate(r); err != nil {
		return err
	}
	return nil
}

type CreateUserRequest struct {
	Name string `validate:"required"`
	Role int    `validate:"required"`
}

func NewCreateUserRequest(name string, role int) *CreateUserRequest {
	return &CreateUserRequest{Name: name, Role: role}
}

func (r *CreateUserRequest) Validate() error {
	if err := validation.Validate(r); err != nil {
		return err
	}
	return nil
}
