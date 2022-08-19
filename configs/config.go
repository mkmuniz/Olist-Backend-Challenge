package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string `json:"port"`
}

type DBConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
	Name string `json:"name"`
}

func init() {
	viper.SetDefault("api.port", "8080")
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", "5432")
}

func Load() {
	viper.SetConfigName("config")
	viper.SetConfigType("to")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Sprintln("Error while loading configs: ", err)
	}
}
