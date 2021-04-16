package config

import (
	"strings"

	"github.com/spf13/viper"
)

/*
	config.go loads and serves configuration parameters using viper.
	It was placed in a seperate file in order to encapsulate the use of viper
	in the event that we would want to use a different library or method in
	the future.
*/
const (
	configFileName      = "config"
	configPathDelimiter = ":"
	configPaths         = "." // Each path is seperated by a `configPathDelimiter`

	productionEnvironmentConfigValue = "prod"
)

func ReadInConfig() error {
	viper.SetConfigName(configFileName)
	for _, configPath := range strings.Split(configPaths, configPathDelimiter) {
		viper.AddConfigPath(configPath)
	}
	return viper.ReadInConfig()
}

func GetString(key string) string {
	return viper.GetString(key)
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
