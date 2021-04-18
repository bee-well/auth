package config

import (
	"os"

	"github.com/spf13/viper"
)

/*
	config.go loads and serves configuration parameters using viper.
	It was placed in a seperate file in order to encapsulate the use of viper
	in the event that we would want to use a different library or method in
	the future.
*/
const (
	productionEnvironmentConfigValue = "prod"
)

func ReadInConfig(configFileName string) error {
	viper.SetConfigName(configFileName)
	return viper.ReadInConfig()
}

func GetString(key string) string {
	if os.Getenv(key) == "" {
		return viper.GetString(key)
	}
	return os.Getenv(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func IsProduction() bool {
	return GetString(AppEnvironment) == productionEnvironmentConfigValue
}
