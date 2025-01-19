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

	// Iterate through the device ids
	for _, deviceId := range deviceIds {
		// Using the array since cursor decode requires an array as parameter
		var tempLocation []model.Location
		// Get the location
		// Specify the query field
		filter := bson.D{{"deviceId", deviceId}}
		// Specify the query option
		opts := options.Find().SetSort(bson.D{{"datePublished", -1}}).SetLimit(1)
		// Query the location
		cursor, err := collection.Find(context.Background(), filter, opts)
		if err != nil {
			log.Err(err).Msg("[Database] Failed to query the locations")
			return []model.Location{}, err
		}

		// decodes the query result
		err = cursor.All(context.Background(), &tempLocation)
		if err != nil {
			log.Err(err).Msg("[Database] Failed to decode location")
			return []model.Location{}, err
		}

		// Append the location into the location array
		locations = append(locations, tempLocation[0])
	}

	return locations, nil
}
