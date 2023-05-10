package di

import (
	"go-patterns/connection"
	"log"
)

func GetUserControllerInstance() *DIController {
	log.Println("Initializing GetUserControllerInstance")

	mongoClient := connection.GetMongoClient()
	repo := NewDIRepository(mongoClient)
	svc := NewDIService(repo)
	ctrl := NewDIController(svc)

	log.Println("Initialized GetUserControllerInstance")

	return ctrl
}

func CreateUserControllerInstance() *DIController {
	log.Println("Initializing CreateUserControllerInstance")

	mongoClient := connection.GetMongoClient()
	repo := NewDIRepository(mongoClient)
	// svc might need different repos while creating user
	svc := NewDIService(repo)
	ctrl := NewDIController(svc)

	log.Println("Initialized CreateUserControllerInstance")

	return ctrl
}
