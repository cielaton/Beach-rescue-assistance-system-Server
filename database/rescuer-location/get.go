package rescuer_location

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"server/model"
)

func GetRescuerLocation(rescuerIds []string, database *mongo.Client) ([]model.RescuerLocation, error) {
	var rescuerLocations []model.RescuerLocation
	// Get the collection
	collection := database.Database("Beach-Rescue-Assistance-System").Collection("Rescuer_Location")

	// Iterate through the rescuer ids
	for _, rescuerId := range rescuerIds {
		// Using the array since cursor decode requires an array as parameter
		var tempLocation []model.RescuerLocation
		// Get the location
		// Specify the query field
		filter := bson.D{{"rescuerId", rescuerId}}
		// Specify the query option
		opts := options.Find().SetSort(bson.D{{"datePublished", -1}}).SetLimit(1)
		// Query the location
		cursor, err := collection.Find(context.Background(), filter, opts)
		if err != nil {
			log.Err(err).Msg("[Database] Failed to query the rescuer locations")
			return []model.RescuerLocation{}, err
		}

		// decodes the query result
		err = cursor.All(context.Background(), &tempLocation)
		if err != nil {
			log.Err(err).Msg("[Database] Failed to decode rescuer location")
			return []model.RescuerLocation{}, err
		}

		// Append the location into the location array
		rescuerLocations = append(rescuerLocations, tempLocation[0])
	}

	return rescuerLocations, nil
}
