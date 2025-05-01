package utils

import "os"

type Environment struct {
	RedisURL      string
	RedisPassword string
	RedisUsername string
	DbHost        string
	DbPassword    string
	DbPort        string
	DbName        string
	DbUsername    string
	Port          string
}

func NewEnv() *Environment {
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	port := os.Getenv("PORT")

	redisURL := os.Getenv("REDIS_URL")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisUsername := os.Getenv("DB_USERNAME")

	return &Environment{
		DbHost:        dbHost,
		DbPassword:    dbPassword,
		DbPort:        dbPort,
		DbName:        dbName,
		DbUsername:    dbUsername,
		Port:          port,
		RedisURL:      redisURL,
		RedisPassword: redisPassword,
		RedisUsername: redisUsername,
	}
}
