package singleton

import (
	"context"
	"go-patterns/connection"
	"go-patterns/model"
	"log"
	"os"

	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SingletonRepository struct{}

func newSingletonRepository() *SingletonRepository {
	log.Println("Initialized Singleton Repository")
	return &SingletonRepository{}
}

func (r *SingletonRepository) getCollection() *mongo.Collection {
	client := connection.GetMongoClient()
	return client.Database(os.Getenv("TEST_DB")).Collection("users")
}

func (r *SingletonRepository) FindByEmail(ctx context.Context, email string) (*model.UserType, error) {
	var user model.UserType
	err := r.getCollection().FindOne(ctx, bson.M{"emailId": email}).Decode(&user)
	return &user, err
}

func (r *SingletonRepository) Create(ctx context.Context, user model.UserType) error {
	_, err := r.getCollection().InsertOne(ctx, user)
	return err
}

var SingletonRepositoryInstance = newSingletonRepository()
