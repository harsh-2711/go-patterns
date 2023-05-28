package service

import (
	"context"
	"go-patterns/di/repository"
	"go-patterns/model"
	"log"
)

type DIService struct {
	repo *repository.DIRepository
}

func NewDIService(repo *repository.DIRepository) *DIService {
	log.Println("Initialized DI Service")
	return &DIService{repo: repo}
}

func (s *DIService) GetUser(ctx context.Context, email string) (*model.UserType, error) {
	return s.repo.FindByEmail(ctx, email)
}

func (s *DIService) CreateUser(ctx context.Context, user model.UserType) error {
	return s.repo.Create(ctx, user)
}
