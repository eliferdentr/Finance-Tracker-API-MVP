package config

import (
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

//the file that loads the configuration from .env file and environment variables

type Config struct {
	AppEnv  string
	AppPort string // viper bind için string tutmak daha pratik, start ederken int’e çevrilir

	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	RedisAddr     string
	RedisPassword string
	RedisDB       int

	JWTSecret     string
	JWTExpiresMin int
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		slog.Warn("No .env file found, reading from environment variables")
	}
	viper.AutomaticEnv()
	viper.BindEnv("DB_HOST")
	viper.BindEnv("DB_PORT")
	viper.BindEnv("DB_USER")
	viper.BindEnv("DB_PASSWORD")
	viper.BindEnv("DB_NAME")
	viper.BindEnv("DB_SSLMODE")
	viper.BindEnv("REDIS_ADDR")
	viper.BindEnv("REDIS_PASSWORD")
	viper.BindEnv("REDIS_DB")
	viper.BindEnv("JWT_SECRET")
	viper.BindEnv("JWT_EXPIRES_MIN")

	// Read values from viper and map to struct
	cfg := &Config{
		AppEnv:        viper.GetString("APP_ENV"),
		AppPort:       viper.GetString("APP_PORT"),
		DBHost:        viper.GetString("DB_HOST"),
		DBPort:        viper.GetInt("DB_PORT"),
		DBUser:        viper.GetString("DB_USER"),
		DBPassword:    viper.GetString("DB_PASSWORD"),
		DBName:        viper.GetString("DB_NAME"),
		DBSSLMode:     viper.GetString("DB_SSLMODE"),
		RedisAddr:     viper.GetString("REDIS_ADDR"),
		RedisPassword: viper.GetString("REDIS_PASSWORD"),
		RedisDB:       viper.GetInt("REDIS_DB"),
		JWTSecret:     viper.GetString("JWT_SECRET"),
		JWTExpiresMin: viper.GetInt("JWT_EXPIRES_MIN"),
	}

	slog.Info("Configuration loaded successfully")
	return cfg, nil
}
