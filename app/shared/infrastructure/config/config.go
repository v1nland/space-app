package config

import (
	"time"

	"github.com/spf13/viper"
)

var (
	// Config is the global configuration object
	Values *Config
)

type (
	// Config -.
	Config struct {
		App      `yaml:"app"`
		Jwt      `yaml:"jwt"`
		Log      `yaml:"logger"`
		Postgres `yaml:"postgres"`
	}

	// App -.
	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
		Prefix  string `yaml:"prefix"`
		Release bool   `yaml:"release"`
	}

	// JWT -.
	Jwt struct {
		Issuer   string `yaml:"issuer"`
		Key      string `yaml:"key"`
		Duration int    `yaml:"duration"`
	}

	// Log -.
	Log struct {
		Level string `yaml:"level"`
	}

	// POSTGRES -.
	Postgres struct {
		Database string `yaml:"database"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Ssl      bool   `yaml:"ssl"`
	}
)

func LoadSettings() {
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./app/shared/infrastructure/config")

	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// Unmarshal config into struct
	if err := viper.Unmarshal(&Values); err != nil {
		panic(err)
	}
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
