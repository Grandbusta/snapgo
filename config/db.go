package config

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewNoSqlDB() *mongo.Client {
	envs := GetEnvs()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(envs.MongoDBURI))
	if err != nil {
		panic(err)
	}
	logrus.WithFields(logrus.Fields{}).Info("Database connected")
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	return client
}
