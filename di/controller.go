package di

import (
	"encoding/json"
	"log"
	"net/http"

	"go-patterns/model"
)

type DIController struct {
	svc *DIService
}

func NewDIController(svc *DIService) *DIController {
	log.Println("Initialized DI Controller")
	return &DIController{svc: svc}
}

func (c *DIController) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	log.Println("Finding user with emailId: ", email)

	user, err := c.svc.GetUser(r.Context(), email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (c *DIController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
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
