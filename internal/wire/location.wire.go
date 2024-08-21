//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/tuanchill/lofola-api/internal/controller"
	"github.com/tuanchill/lofola-api/internal/repo"
	"github.com/tuanchill/lofola-api/internal/service"
)

func InitLocationRouterHandler() (*controller.LocationController, error) {
	wire.Build(
		repo.NewLocationRepo,
		service.NewLocationService,
		controller.NewLocationController,
	)
	return new(controller.LocationController), nil
}
