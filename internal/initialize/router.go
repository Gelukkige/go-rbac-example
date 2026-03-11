package initialize

import (
	"fmt"
	"go-rbac-example/internal/global"

	"github.com/gin-gonic/gin"
)

func RouterInit() {
	serverConfig := global.Config.Server

	r := gin.Default()
	addr := fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)
	go r.Run(addr)
}
