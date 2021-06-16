package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string
		Host string
	}
	Database struct {
		User   string
		Pass   string
		DbName string
		Port   int
		Host   string
	}
}

func NewConfig() *Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")

	var config Config
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	errC := viper.Unmarshal(&config)

	if errC != nil {
		fmt.Printf("Unable to decode into struct, %v", errC)
	}

	return &config
}
