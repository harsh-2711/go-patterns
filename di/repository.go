package di

import (
	"context"
	"go-patterns/model"
	"log"
	"os"

	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DIRepository struct {
	client *mongo.Client
}

func NewDIRepository(client *mongo.Client) *DIRepository {
	log.Println("Initialized DI Repository")
	return &DIRepository{client: client}
}

func (r *DIRepository) getCollection() *mongo.Collection {
	return r.client.Database(os.Getenv("TEST_DB")).Collection("users")
}

func (r *DIRepository) FindByEmail(ctx context.Context, email string) (*model.UserType, error) {
	var user model.UserType
	err := r.getCollection().FindOne(ctx, bson.M{"emailId": email}).Decode(&user)
	return &user, err
}

func (r *DIRepository) Create(ctx context.Context, user model.UserType) error {
	_, err := r.getCollection().InsertOne(ctx, user)
	return err
}
