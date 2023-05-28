package route

import (
	"go-patterns/di/controller"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewHandler),
	fx.Invoke(registerRoutes),
)

type Handler struct {
	Gin *gin.Engine
}

func NewHandler() Handler {
	engine := gin.Default()
	return Handler{Gin: engine}
}

func registerRoutes(
	userController *controller.DIController,
	handler Handler,
) {
	userRoutes := handler.Gin.Group("/di/user")
	{
		userRoutes.GET("/", userController.GetUserHandler())
		userRoutes.POST("/create", userController.CreateUserHandler())
	}
}