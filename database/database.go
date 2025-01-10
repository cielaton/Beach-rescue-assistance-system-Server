package database

import (
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"os"
)

func Connect() (*mongo.Client, error) {
	databaseClient, err := mongo.Connect(options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	if err != nil {
		log.Err(err).Msg("[Database] Failed to connect to the database")
		return nil, err
	}

	log.Info().Msg("[Database] Connected to the database")

	return databaseClient, nil
}
