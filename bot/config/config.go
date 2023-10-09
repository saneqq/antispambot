package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	BotToken string `mapstructure:"BOT_TOKEN"`
	BotName  string `mapstructure:"BOT_NAME"`
}

var config *Config

func GetConfig() *Config {
	return config
}

func init() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read config file")
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to unmarshal config file")
	}
}
