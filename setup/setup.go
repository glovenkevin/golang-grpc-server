package setup

import "github.com/spf13/viper"

func Load() error {
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
