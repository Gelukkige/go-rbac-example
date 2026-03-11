package global

import (
	"go-rbac-example/internal/config"

	"gorm.io/gorm"
)

var (
	Config config.Config
	DB     *gorm.DB
)
