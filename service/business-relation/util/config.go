package util

import (
	"github.com/calvinkmts/expert-pancake/engine/config"
	"github.com/spf13/viper"
)

type Config struct {
	Db     config.PostgresqlConfig `mapstructure:"db"`
	Server config.ServerConfig     `mapstructure:"server"`
}

var vp *viper.Viper

func LoadConfig() (config Config, err error) {
	vp = viper.New()

	vp.SetConfigName("business_relation")
	vp.SetConfigType("json")
	vp.AddConfigPath("../../config")

	err = vp.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	err = vp.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
