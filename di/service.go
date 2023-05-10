package di

import (
	"context"
	"go-patterns/model"
	"log"
)

type DIService struct {
	repo *DIRepository
}

func NewDIService(repo *DIRepository) *DIService {
	log.Println("Initialized DI Service")
	return &DIService{repo: repo}
}

func (s *DIService) GetUser(ctx context.Context, email string) (*model.UserType, error) {
	return s.repo.FindByEmail(ctx, email)
}

func (s *DIService) CreateUser(ctx context.Context, user model.UserType) error {
	return s.repo.Create(ctx, user)
}
