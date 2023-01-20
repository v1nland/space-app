package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/viper"
)

const DURATION = time.Second * 5

func LoadSettings(squadName string, appName string) {
	if err := getConfig(); err != nil {
		panic("no config found")
	}

	if len(appName) == 0 {
		panic("no app to run")
	}

	viper.Set("app", appName)
}

func Get(key string) interface{} {
	return viper.Get(key)
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetInt64(key string) int64 {
	return viper.GetInt64(key)
}

func GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

func GetDuration(key string) time.Duration {
	return viper.GetDuration(key)
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func AllSettings() map[string]interface{} {
	return viper.AllSettings()
}

func getConfig() error {
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./app/shared/infrastructure/config")

	err := viper.ReadInConfig()
	if err != nil {
		return errors.New(fmt.Sprintf("fatal error config file: %s \n", err))
	}

	return nil
}
