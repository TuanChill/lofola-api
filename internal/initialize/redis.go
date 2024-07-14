package initialize

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/tuanchill/lofola-api/global"
)

var ctx = context.Background()

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", global.Config.Redis.Host, global.Config.Redis.Port),
		Password: global.Config.Redis.Password, // no password set
		DB:       global.Config.Redis.DB,       // use default DB
		PoolSize: 10,
	})

	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		_, err := rdb.Ping(ctx).Result()
		if err != nil {
			fmt.Println("Redis connection failed, retrying...")
			time.Sleep(5 * time.Second)
		} else {
			fmt.Println("Redis connection successful")
			global.MRB = rdb
			return
		}
	}

	panic("failed to connect redis")
}
