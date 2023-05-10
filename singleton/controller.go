package singleton

import (
	"encoding/json"
	"log"
	"net/http"

	"go-patterns/model"
)

type SingletonController struct {
	svc Service
}

func newSingletonController(svc Service) *SingletonController {
	log.Println("Initialized Singleton Controller")
	return &SingletonController{svc: svc}
}

func (c *SingletonController) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	log.Println("Finding user with emailId: ", email)

	user, err := c.svc.GetUser(r.Context(), email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (c *SingletonController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.UserType
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("Creating user with emailId: ", user.EmailID)

	err = c.svc.CreateUser(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

var SingletonControllerInstance = newSingletonController(SingletonServiceInstance)
