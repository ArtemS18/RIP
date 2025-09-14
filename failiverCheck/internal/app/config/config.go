package config

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	ServiceHost string `mapstructure:"ServiceHost"`
	ServicePort int    `mapstructure:"ServicePort"`
}

func NewConfig() (*Config, error) {
	configName := "config"
	_ = godotenv.Load()
	if envName := os.Getenv("CONFIG_NAME"); envName != "" {
		configName = envName
	}

	viper.SetConfigName(configName)
	viper.AddConfigPath("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	newConf := &Config{}
	err = viper.Unmarshal(newConf)
	if err != nil {
		return nil, err
	}
	log.Info("Config created")
	return newConf, nil
}
