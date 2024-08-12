//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/tuanchill/lofola-api/internal/controller"
	"github.com/tuanchill/lofola-api/internal/repo"
	"github.com/tuanchill/lofola-api/internal/service"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepo,
		service.NewUserService,
		controller.NewUserController,
	)
	return new(controller.UserController), nil

}
