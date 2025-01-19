package user

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"server/model"
)

func GetUser(userName string, database *mongo.Client) (model.User, error) {
	user := model.User{}

	// Get the collection
	collection := database.Database("Beach-Rescue-Assistance-System").Collection("User")

	// Get the user
	// Query by user ID
	filter := bson.D{{"userId", userName}}
	err := collection.FindOne(context.Background(), filter).Decode(&user)

	if err != nil {
		log.Err(err).Msg("[Database] Failed to get the user info")
		return model.User{}, err
	}

	return user, nil
}
