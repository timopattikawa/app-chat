package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	SupabaseCon Supabase
	DatabaseCon Database
	ServerCon   Server
}

type Database struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

type Server struct {
	Host string
	Port string
}

type Supabase struct {
	Url string
	Key string
}

func GetConfig() *Config {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Fail to read config file")
		return nil
	}
	return &Config{
		Supabase{
			Url: viper.GetString("URL_SUPABASE"),
			Key: viper.GetString("API_KEY_SUPABASE"),
		},
		Database{
			Username: viper.GetString(`database.username`),
			Password: viper.GetString(`database.pass`),
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
