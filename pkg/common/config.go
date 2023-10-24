package common

import "github.com/spf13/viper"

func LoadConfig() error {
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	return viper.ReadInConfig()
}

func GetConfig(key string) string {
	return viper.GetString(key)
}
