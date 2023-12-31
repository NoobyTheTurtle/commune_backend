package config

import (
	"github.com/spf13/viper"
	"gorm.io/gorm/logger"
)

type Config struct {
	ServerPort string          `mapstructure:"SERVER_PORT"`
	DBPort     string          `mapstructure:"DB_PORT"`
	DBName     string          `mapstructure:"DB_NAME"`
	DBPassword string          `mapstructure:"POSTGRES_PASSWORD"`
	GinMode    string          `mapstructure:"GIN_MODE"`
	DBLogLevel logger.LogLevel `mapstructure:"DB_LOG_LEVEL"`
}

func Load() (*Config, error) {
	c := Config{}
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
