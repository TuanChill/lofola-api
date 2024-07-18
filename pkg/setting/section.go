package setting

type Config struct {
	// MongoDB MongoDbSetting
	Mysql    MySQLSetting `mapstructure:"mysql"`
	Redis    RedisSetting `mapstructure:"redis"`
	Server   ServerSetting
	Security SecurityConfig `mapstructure:"security"`
}

type MongoDbSetting struct {
	URI string `mapstructure:"uri"`
}

type MySQLSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	UserName string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"database"`
	MaxIdle  int    `mapstructure:"maxIdleConns"`
	MaxOpen  int    `mapstructure:"maxOpenConns"`
	MaxLife  int    `mapstructure:"maxConnLifetime"`
}

type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type ServerSetting struct {
	Port int `mapstructure:"port"`
}

type SecurityConfig struct {
	AccessTokenSecret  SecretKey `mapstructure:"accessTokenSecret"`
	RefreshTokenSecret SecretKey `mapstructure:"refreshTokenSecret"`
}

type SecretKey struct {
	SecretKey string `mapstructure:"secret"`
}
