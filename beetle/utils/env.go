package utils

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
	return &Environment{}
}
