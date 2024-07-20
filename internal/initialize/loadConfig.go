package initialize

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/tuanchill/lofola-api/global"
)

func LoadConfig(path string) {
	viper.SetConfigName("config")
	viper.AddConfigPath(path)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}
