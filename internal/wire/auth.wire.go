//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/tuanchill/lofola-api/internal/controller"
	"github.com/tuanchill/lofola-api/internal/repo"
	"github.com/tuanchill/lofola-api/internal/repo/redis"
	"github.com/tuanchill/lofola-api/internal/service"
)

func InitAuthRouterHandler() (*controller.AuthController, error) {
	wire.Build(
		repo.NewUserRepo,
		service.NewAuthService,
		redis.NewOtpRedisRepo,
		redis.NewTokenRedisRepo,
		controller.NewAuthController,
	)
	return new(controller.AuthController), nil

}
