package device

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func DeleteDevice(deviceId string, database *mongo.Client) error {
	// Get the collection
	collection := database.Database("Beach-Rescue-Assistance-System").Collection("Device")

	// Delete the device
	// Specify the filter
	filter := bson.M{"deviceId": deviceId}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Err(err).Msg("[Device] Failed to delete device")
		return err
	}

	return nil
}
