package config

import (
	"os"
)

type Envs struct {
	MongoDBURI string
	JWT_SECRET string
	PORT       string
}

func GetEnv(key string, default_value string) string {
	val := os.Getenv(key)
	if val == "" {
		return default_value
	}
	return val
}

func GetEnvs() *Envs {
	return &Envs{
		MongoDBURI: GetEnv("MONGODB_URI", ""),
		JWT_SECRET: GetEnv("JWT_SECRET", ""),
		PORT:       GetEnv("PORT", ":3000"),
	}
}
