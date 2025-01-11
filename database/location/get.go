package location

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"server/model"
)

func GetLocation(database *mongo.Client) (model.Location, error) {
	var location = model.Location{}
	// Get the collection
	collection := database.Database("Beach-Rescue-Assistance-System").Collection("Location")

	// Get the location
	// Specify the query field
	filter := bson.D{{"deviceId", "abcd"}}
	err := collection.FindOne(context.Background(), filter).Decode(&location)
	if err != nil {
		log.Err(err).Msg("[Database] Failed to get location info")
		return model.Location{}, err
	}

	return location, nil
}
