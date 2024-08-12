//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/tuanchill/lofola-api/internal/controller"
	"github.com/tuanchill/lofola-api/internal/repo"
	"github.com/tuanchill/lofola-api/internal/service"
)

func InitGroupRouterHandler() (*controller.GroupController, error) {
	wire.Build(
		repo.NewGroupRepo,
		repo.NewUserRepo,
		service.NewGroupService,
		controller.NewGroupController,
	)
	return new(controller.GroupController), nil
}
