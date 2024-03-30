package model

import (
	"time"

	"github.com/boon-neko/golang-grpc-backend-example/internal/pkg/uuid"
)

type Book struct {
	ID        string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Books []*Book

func NewBook(title string) *Book {
	return &Book{ID: uuid.New(), Title: title}
}
