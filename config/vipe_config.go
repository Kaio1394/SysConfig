package config

import (
	"github.com/spf13/viper"
)

var ConfigViper Config

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		return
	}
	var configs Config
	if err := viper.Unmarshal(&configs); err != nil {
		return
	}
	ConfigViper = configs
}
