package initialize

func Init(configFile string) {
	LoadConfig(configFile)
	DBInit()
	RouterInit()
}
