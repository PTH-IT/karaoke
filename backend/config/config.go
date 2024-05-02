package config

import (
	"encoding/json"
	configdev "karaoke/config/dev"
	configLocal "karaoke/config/local"
	configprod "karaoke/config/production"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Env       string `json:"env"`
	Port      string `json:"port"`
	Aws       AwsConfig
	Redis     RedisConfig `json:"redis"`
	PostgreDB PostgreDB
}
type PostgreDB struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Pass string `json:"password"`
	Db   string `json:"db"`
}
type RedisConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Pass     string `json:"password"`
	Username string `json:"Username"`
	Db       string `json:"db"`
}
type AwsConfig struct {
	Host   string
	Port   string
	Region string
	Id     string
	Secret string
	Token  string
}

func Getconfig() AppConfig {
	var appConfig AppConfig
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	env := os.Getenv("ENVIRONMENT")
	if env == "local" {
		err := json.Unmarshal([]byte(configLocal.ConfigApp), &appConfig)
		if err != nil {
			panic(err)
		}
	} else if env == "dev" {
		err := json.Unmarshal([]byte(configdev.ConfigApp), &appConfig)
		if err != nil {
			panic(err)
		}
	} else if env == "production" {
		err := json.Unmarshal([]byte(configprod.ConfigApp), &appConfig)
		if err != nil {
			panic(err)
		}
	}
	appConfig.Aws.Host = os.Getenv("AWS_HOST")
	appConfig.Aws.Port = os.Getenv("AWS_PORT")
	appConfig.Aws.Region = os.Getenv("AWS_REGION")
	appConfig.Aws.Id = os.Getenv("AWS_ID")
	appConfig.Aws.Secret = os.Getenv("AWS_SECRET")
	appConfig.Aws.Token = os.Getenv("AWS_TOKEN")
	if appConfig.Port == "" {
		appConfig.Port = os.Getenv("PORT")
	}
	//MYSQL

	return appConfig
}
