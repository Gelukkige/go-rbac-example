package initialize

import (
	"go-rbac-example/internal/permission"
	"go-rbac-example/internal/schema"
)

func Init(configFile string) {
	schema.Init()
	LoadConfig(configFile)
	DBInit()
	RedisInit()
	permission.Init()
	RouterInit()
}
