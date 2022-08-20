package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

var cfg *Config

func init() {
	viper.SetDefault("api.port", "8080")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Sprintln("Error while loading configs: ", err)
	}

	cfg = new(Config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}
	cfg.DB = DBConfig{
		URL:  viper.GetString("db.url"),
		Name: viper.GetString("db.name"),
	}

	return nil
}

func GetAPIConfig() string {
	return cfg.API.Port
}

func GetDBConfig() DBConfig {
	return cfg.DB
}
