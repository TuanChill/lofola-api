package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/tuanchill/lofola-api/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	MDB    *gorm.DB
	RDB    *redis.Client
)
