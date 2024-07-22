package redis

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/tuanchill/lofola-api/configs/common/constants"
	"github.com/tuanchill/lofola-api/pkg/utils"
)

func saveTokenBlack(ctx *gin.Context, r *redis.Client, key string, expireTime int) error {
	err := r.Set(ctx, key, "blacklisted", time.Duration(time.Duration(expireTime).Seconds())).Err()

	if err != nil {
		return err
	}

	return nil
}

func SaveAccessTokenBlack(ctx *gin.Context, r *redis.Client, token string) error {
	key := utils.FormatKeyRedis(constants.AccessTokenBlack, token)

	return saveTokenBlack(ctx, r, key, constants.ExpiresAccessToken)
}

func SaveRefreshTokenBlack(ctx *gin.Context, r *redis.Client, token string) error {
	key := utils.FormatKeyRedis(constants.RefreshTokenBlack, token)

	return saveTokenBlack(ctx, r, key, constants.ExpiresRefreshToken)
}

func IsTokenBlack(ctx *gin.Context, r *redis.Client, key string) (bool, error) {
	_, err := r.Get(ctx, key).Result()

	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}
