package global

import (
	"go-rbac-example/internal/config"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	Config      config.Config
	DB          *gorm.DB
	RedisClient *redis.Client
)
