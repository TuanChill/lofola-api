package redis

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/tuanchill/lofola-api/configs/common/constants"
	"github.com/tuanchill/lofola-api/internal/models"
)

func SpamUser(ctx *gin.Context, rdb *redis.Client, email string, requestThreshold int64) *models.SpamUserRedisResponse {
	key := fmt.Sprintf("%s_%s", constants.SpamKey, email)

	numberRequest, err := rdb.Incr(ctx, key).Result()
	if err != nil {
		return nil
	}

	var ttl time.Duration

	// check if the number of requests exceeds the threshold
	if numberRequest == requestThreshold+1 {
		rdb.Expire(ctx, key, constants.InitialBlock)
		ttl = constants.InitialBlock
	} else if numberRequest > requestThreshold+2 {
		rdb.Expire(ctx, key, constants.ExtendBlock)
		ttl = constants.ExtendBlock
	} else {
		ttl, err = rdb.TTL(ctx, key).Result()
		if err != nil {
			return nil
		}
	}

	// return ttl if the number of requests exceeds the threshold
	if numberRequest > requestThreshold {
		return &models.SpamUserRedisResponse{
			ExpireTime: int(ttl.Seconds()),
			IsSpam:     true,
		}
	}

	rdb.Expire(ctx, key, constants.ExpireDuration)

	return &models.SpamUserRedisResponse{
		ExpireTime: 0,
		IsSpam:     false,
	}
}
