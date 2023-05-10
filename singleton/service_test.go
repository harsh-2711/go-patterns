package singleton

import (
	"context"
	"go-patterns/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestSingletonService_GetUser(t *testing.T) {
	name := "Test User"
	emailId := "test@example.com"

	mockRepo := &MockRepository{
		users: map[string]*model.UserType{
			emailId: {Name: &name, EmailID: &emailId},
		},
	}
	service := newSingletonService(mockRepo)

	ctx := context.Background()
	user, err := service.GetUser(ctx, emailId)
	assert.NoError(t, err)
	assert.Equal(t, name, *user.Name)
	assert.Equal(t, emailId, *user.EmailID)

	_, err = service.GetUser(ctx, "nonexistent@example.com")
	assert.Error(t, err)
}

func TestSingletonService_CreateUser(t *testing.T) {
	id := primitive.NewObjectID().String()
	name := "Test User"
	emailId := "new@example.com"
	lastLogin := primitive.NewDateTimeFromTime(time.Now())

	mockRepo := &MockRepository{users: map[string]*model.UserType{}}
	service := newSingletonService(mockRepo)

	ctx := context.Background()
	newUser := model.UserType{ID: id, Name: &name, EmailID: &emailId, LastLogin: &lastLogin}
	err := service.CreateUser(ctx, newUser)
	assert.NoError(t, err)

	user, err := mockRepo.FindByEmail(ctx, emailId)
	assert.NoError(t, err)
	assert.Equal(t, id, user.ID)
	assert.Equal(t, name, *user.Name)
	assert.Equal(t, emailId, *user.EmailID)
	assert.Equal(t, lastLogin, *user.LastLogin)
}
