package configs

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	App      App
	Database Database
	Jwt      Jwt
}

type Database struct {
	Host            string
	Port            string
	User            string
	Password        string
	DbName          string
	Debug           bool
	MaxIdleConns    int
	MaxOpenConns    int
	MaxConnLifeTime int
	MaxConnIdleTime int
}

type App struct {
	Host  string
	Port  string
	Name  string
	Env   string
	Debug bool
}

type Jwt struct {
	Secret string
}

var Configs *Config

func Load(path string) *Config {
	err := godotenv.Load(path + "/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	v := viper.New()
	v.AutomaticEnv()

	config := Config{
		Database: Database{
			Host:            v.GetString("DB_HOST"),
			Port:            v.GetString("DB_PORT"),
			User:            v.GetString("DB_USER"),
			Password:        v.GetString("DB_PASSWORD"),
			DbName:          v.GetString("DB_NAME"),
			Debug:           v.GetBool("DB_DEBUG"),
			MaxIdleConns:    v.GetInt("DB_MAX_IDLE_CONNS"),
			MaxOpenConns:    v.GetInt("DB_MAX_OPEN_CONNS"),
			MaxConnLifeTime: v.GetInt("DB_MAX_CONN_LIFE_TIME"),
			MaxConnIdleTime: v.GetInt("DB_MAX_CONN_IDLE_TIME"),
		},
		App: App{
			Host:  v.GetString("APP_HOST"),
			Port:  v.GetString("APP_PORT"),
			Name:  v.GetString("APP_NAME"),
			Env:   v.GetString("APP_ENV"),
			Debug: v.GetBool("APP_DEBUG"),
		},
		Jwt: Jwt{
			Secret: v.GetString("JWT_SECRET"),
		},
	}
	Configs = &config
	return &config
}
