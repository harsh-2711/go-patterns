package bundler

import (
	"context"
	"fmt"
	"go-patterns/connection"
	"go-patterns/di/controller"
	"go-patterns/di/repository"
	"go-patterns/di/route"
	"go-patterns/di/service"
	"go-patterns/di/storage"
	"log"

	"go.uber.org/fx"
)

var Module = fx.Options(
	controller.Module,
	service.Module,
	repository.Module,
	storage.Module,
	route.Module,
	fx.Invoke(runApp),
)

func runApp(lc fx.Lifecycle, handler route.Handler) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			log.Println("Starting DI pattern")
			connection.ConnectMongo()
			go handler.Gin.Run(":9090")
			return nil
		},
		OnStop: func(context.Context) error {
			fmt.Println("Stopping Application")
			return nil
		},
	})
}