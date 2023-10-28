package config

import (
	"github.com/spf13/viper"
	"log"
)

func Get() *Config {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Fail to read config file", err)
		return nil
	}

	return &Config{
		Supabase{
			Url: viper.GetString("URL_SUPABASE"),
			Key: viper.GetString("API_KEY_SUPABASE"),
		},
		Database{
			Username: viper.GetString(`database.username`),
			Password: viper.GetString(`database.password`),
			Host:     viper.GetString(`database.host`),
			Port:     viper.GetString(`database.port`),
			Name:     viper.GetString(`database.name`),
		},
		Server{
			Host: viper.GetString(`server.host`),
			Port: viper.GetString(`server.port`),
		},
	}
}
