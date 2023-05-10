package singleton

import (
	"context"
	"go-patterns/model"
)

type Repository interface {
	FindByEmail(ctx context.Context, email string) (*model.UserType, error)
	Create(ctx context.Context, user model.UserType) error
}

type Service interface {
	GetUser(ctx context.Context, email string) (*model.UserType, error)
	CreateUser(ctx context.Context, user model.UserType) error
}
