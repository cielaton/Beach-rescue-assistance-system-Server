package rescuer

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"server/model"
)

func GetRescuerBySafeAreaId(safeAreaId string, database *mongo.Client) ([]model.Rescuer, error) {
	var rescuers []model.Rescuer
	// Get the collection
	collection := database.Database("Beach-Rescue-Assistance-System").Collection("Rescuer")

	// Get the rescuer
	// Specify the query field
	filter := bson.D{{"safeAreaId", safeAreaId}}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Err(err).Msg("[Database] Failed to get rescuer info")
		return []model.Rescuer{}, err
	}
	// Decode all the results
	err = cursor.All(context.Background(), &rescuers)
	if err != nil {
		log.Err(err).Msg("[Database] Failed to decode rescuer info")
		return []model.Rescuer{}, err
	}

	return rescuers, nil
}

func GetRescuerByRescuerId(rescuerId string, database *mongo.Client) (model.Rescuer, error) {
	var rescuer = model.Rescuer{}
	// Get the collection
	collection := database.Database("Beach-Rescue-Assistance-System").Collection("Rescuer")

	// Get the rescuer
	// Specify the query field
	filter := bson.D{{"rescuerId", rescuerId}}
	err := collection.FindOne(context.Background(), filter).Decode(&rescuer)
	if err != nil {
		log.Err(err).Msg("[Database] Failed to get rescuer info")
		return model.Rescuer{}, err
	}

	return rescuer, nil
}
