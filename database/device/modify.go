package device

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func ChangeDeviceActiveStatus(deviceId string, status bool, database *mongo.Client) error {
	collection := database.Database("Beach-Rescue-Assistance-System").Collection("Device")

	// Modify the device
	// Specify the filter
	filter := bson.D{{"deviceId", deviceId}}
	// Specify the update
	update := bson.D{{"$set", bson.D{{"isEnabled", status}}}}
	// Update the field
	result, err := collection.UpdateOne(context.Background(), filter, update)

	if result.MatchedCount == 0 {
		err = errors.New("unmatched device id")
	}
	if err != nil {
		log.Err(err).Msg("[Device] Failed to change active status")
		return err
	}

	return nil
}
