package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	"server/database/location"
	"server/model"
)

func GetLocationHandler(echoContext echo.Context, database *mongo.Client) error {
	// Bind the POST request body
	var requestBody model.LocationRequest
	err := echoContext.Bind(&requestBody)
	if err != nil {
		log.Err(err).Msg("[Location] Invalid request body")
		return echoContext.JSON(http.StatusBadRequest,
			map[string]any{"error": err.Error()})
	}

	// Get the locations
	locations, err := location.GetLocation(requestBody.DeviceIds, database)
	if err != nil {
		log.Err(err).Msg("[Database] Failed to get locations ")
		return echoContext.JSON(http.StatusInternalServerError,
			map[string]any{"error": err.Error()})
	}

	return echoContext.JSON(http.StatusOK, locations)
}
