package singleton

import (
	"go-patterns/connection"
	"log"
	"net/http"
)

func Init() {
	log.Println("Starting Singleton pattern")
	connection.ConnectMongo()

	http.HandleFunc("/singleton/user", SingletonControllerInstance.GetUserHandler)
	http.HandleFunc("/singleton/user/create", SingletonControllerInstance.CreateUserHandler)

	http.ListenAndServe(":8080", nil)
}
