package safe_area

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"server/model"
)

func GetSafeArea(database *mongo.Client) (model.SafeArea, error) {
	var safeArea = model.SafeArea{}
	// Get the collection
	collection := database.Database("Beach-Rescue-Assistance-System").Collection("Safe-Area")

	// Get the safeArea
	// Specify the query field
	filter := bson.D{{"safeAreaId", "abcd"}}
	err := collection.FindOne(context.Background(), filter).Decode(&safeArea)
	if err != nil {
		log.Err(err).Msg("[Database] Failed to get safeArea info")
		return model.SafeArea{}, err
	}

	return safeArea, nil
}
