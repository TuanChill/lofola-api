package initialize

func Run() {
	LoadConfig("./configs")

	InitMysql()

	InitRedis()

	InitRouter()
}
