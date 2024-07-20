package initialize

func Run() {
	LoadConfig("./configs/yaml")

	InitMysql()

	InitRedis()

	InitRouter()
}
