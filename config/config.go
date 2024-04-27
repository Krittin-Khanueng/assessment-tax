package config

import (
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type (
	Config struct {
		Server   *Server   `mapstructure:"server" validation:"required"`
		Database *Database `mapstructure:"database" validation:"required"`
	}

	Server struct {
		Port         int           `mapstructure:"port" validation:"required"`
		AllowOrigins []string      `mapstructure:"allowOrigins" validation:"required"`
		BodyLimit    string        `mapstructure:"bodyLimit" validation:"required"`
		TimeOut      time.Duration `mapstructure:"timeout" validation:"required"`
	}


	Database struct {
		Host     string `mapstructure:"host" validate:"required"`
		Port     int    `mapstructure:"port" validate:"required"`
		User     string `mapstructure:"user" validate:"required"`
		Password string `mapstructure:"password" validate:"required"`
		DBName   string `mapstructure:"dbname" validate:"required"`
		SSLMode  string `mapstructure:"sslmode" validate:"required"`
		Schema   string `mapstructure:"schema" validate:"required"`
	}
)

var (
	once           sync.Once
	configInstance *Config
)

func ConfigGetting() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}

		validating := validator.New()

		if err := validating.Struct(configInstance); err != nil {
			panic(err)
		}

	})
	//
	return configInstance
}
