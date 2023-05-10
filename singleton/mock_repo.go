package singleton

import (
	"context"
	"fmt"
	"go-patterns/model"
)

type MockRepository struct {
	users map[string]*model.UserType
}

func (m *MockRepository) FindByEmail(ctx context.Context, email string) (*model.UserType, error) {
	user, ok := m.users[email]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (m *MockRepository) Create(ctx context.Context, user model.UserType) error {
	m.users[*user.EmailID] = &user
	return nil
}
