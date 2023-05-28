package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"go-patterns/di/service"
	"go-patterns/model"

	"github.com/gin-gonic/gin"
)

type DIController struct {
	svc *service.DIService
}

func NewDIController(svc *service.DIService) *DIController {
	log.Println("Initialized DI Controller")
	return &DIController{svc: svc}
}

func (c *DIController) GetUserHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := ctx.Request.URL.Query().Get("email")
		log.Println("Finding user with emailId: ", email)

		user, err := c.svc.GetUser(ctx.Request.Context(), email)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}

func (c *DIController) CreateUserHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.UserType
		err := json.NewDecoder(ctx.Request.Body).Decode(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		log.Println("Creating user with emailId: ", user.EmailID)

		err = c.svc.CreateUser(ctx.Request.Context(), user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusCreated, nil)
	}
}
