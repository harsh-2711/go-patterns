package di

import (
	"go-patterns/di/bundler"

	"go.uber.org/fx"
)

func Init() {
	fx.New(bundler.Module).Run()
}
