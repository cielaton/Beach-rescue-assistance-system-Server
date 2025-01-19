package location

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"server/model"
)

func GetLocation(deviceIds []string, database *mongo.Client) ([]model.Location, error) {
	var locations []model.Location
	// Get the collection
	collection := database.Database("Beach-Rescue-Assistance-System").Collection("Location")

	// Get the location
	// Specify the query field
	filter := bson.D{{"deviceId", bson.M{"$in": deviceIds}}}
	// Specify the query option
	opts := options.Find().SetSort(bson.D{{"datePublished", -1}}).SetLimit(1)
	// Query the location
	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		log.Err(err).Msg("[Database] Failed to query the locations")
		return []model.Location{}, err
	}

	// decodes the query result
	err = cursor.All(context.Background(), &locations)
	if err != nil {
		log.Err(err).Msg("[Database] Failed to decode locations")
		return []model.Location{}, err
	}

	return locations, nil
}
