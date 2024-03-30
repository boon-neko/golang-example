package model

import (
	"time"

	"github.com/boon-neko/golang-grpc-backend-example/internal/pkg/uuid"
)

// User is User struct
type User struct {
	ID        string
	Name      string
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Users []*Users

// UserWithBooks is User struct with Books
type UserWithBooks struct {
	User
	Books
}

type UsersWithBooks []*UserWithBooks

// NewUser is User Constructor
func NewUser(name string, role uint) *User {
	return &User{ID: uuid.New(), Name: name, Role: NewRole(role)}
}
