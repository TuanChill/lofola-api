package initialize

func Run() {
	LoadConfig("./configs/yaml")

	InitValidator()

	InitMysql()

	InitRedis()

	InitRouter()
}
