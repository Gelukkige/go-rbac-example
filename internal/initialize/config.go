package initialize

import (
	"fmt"
	"go-rbac-example/internal/global"

	"github.com/spf13/viper"
)

func LoadConfig(configFile string) {
	v := viper.New()
	v.SetConfigFile(configFile)

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("读取配置文件失败: %s", err))
	}

	if err := v.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("配置反序列化失败: %s", err))
	}
}
