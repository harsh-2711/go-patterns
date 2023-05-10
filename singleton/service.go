package singleton

import (
	"context"
	"go-patterns/model"
	"log"
)

type SingletonService struct {
	repo Repository
}

func newSingletonService(repo Repository) *SingletonService {
	log.Println("Initialized Singleton Service")
	return &SingletonService{repo: repo}
}

func (s *SingletonService) GetUser(ctx context.Context, email string) (*model.UserType, error) {
	return s.repo.FindByEmail(ctx, email)
}

func (s *SingletonService) CreateUser(ctx context.Context, user model.UserType) error {
	return s.repo.Create(ctx, user)
}

var SingletonServiceInstance = newSingletonService(SingletonRepositoryInstance)
