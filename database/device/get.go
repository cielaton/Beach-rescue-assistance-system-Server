package device

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"server/model"
)

func GetDeviceBySafeAreaId(safeAreaId string, database *mongo.Client) ([]model.Device, error) {
	var devices []model.Device
	// Get the collection
	collection := database.Database("Beach-Rescue-Assistance-System").Collection("Device")

	// Get the device
	// Specify the query field
	filter := bson.D{{"safeAreaId", safeAreaId}}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Err(err).Msg("[Database] Failed to get device info")
		return []model.Device{}, err
	}
	// Decode all the results
	err = cursor.All(context.Background(), &devices)
	if err != nil {
		log.Err(err).Msg("[Database] Failed to decode device info")
		return []model.Device{}, err
	}

	return devices, nil
}

func GetDeviceByDeviceId(deviceId string, database *mongo.Client) (model.Device, error) {
	var device = model.Device{}
	// Get the collection
	collection := database.Database("Beach-Rescue-Assistance-System").Collection("Device")

	// Get the device
	// Specify the query field
	filter := bson.D{{"deviceId", deviceId}}
	err := collection.FindOne(context.Background(), filter).Decode(&device)
	if err != nil {
		log.Err(err).Msg("[Database] Failed to get device info")
		return model.Device{}, err
	}

	return device, nil
}
