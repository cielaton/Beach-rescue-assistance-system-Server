package database

import (
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func DatabaseConnect() (*mongo.Client, error) {
	databaseClient, err := mongo.Connect(options.Client().ApplyURI(""))
	if err != nil {
		log.Err(err).Msg("[Database] Failed to connect to the database")
		return nil, err
	}

	log.Info().Msg("[Database] Connected to the database")

	return databaseClient, nil
}
