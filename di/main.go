package di

import (
	"go-patterns/connection"
	"log"
	"net/http"
)

func Init() {
	log.Println("Starting DI pattern")
	connection.ConnectMongo()

	http.HandleFunc("/di/user", func(w http.ResponseWriter, r *http.Request) {
		DIControllerInstance := GetUserControllerInstance()
		DIControllerInstance.GetUserHandler(w, r)
	})

	http.HandleFunc("/di/user/create", func(w http.ResponseWriter, r *http.Request) {
		DIControllerInstance := CreateUserControllerInstance()
		DIControllerInstance.CreateUserHandler(w, r)
	})

	http.ListenAndServe(":8080", nil)
}
