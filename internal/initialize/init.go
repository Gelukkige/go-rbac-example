package initialize

import "go-rbac-example/internal/permission"

func Init(configFile string) {
	LoadConfig(configFile)
	DBInit()
	RedisInit()
	permission.Init()
	RouterInit()
}
