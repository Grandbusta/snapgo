package config

import "os"

type Envs struct {
	MongoDBURI string
	JWT_SECRET string
}

func GetEnvs() *Envs {
	return &Envs{
		MongoDBURI: os.Getenv("MONGODB_URI"),
		JWT_SECRET: os.Getenv("JWT_SECRET"),
	}
}
