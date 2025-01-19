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
		msg := "[Location] Invalid request body"
		log.Err(err).Msg(msg)
		return echoContext.JSON(http.StatusBadRequest,
			map[string]any{
				"message": msg,
				"error":   err.Error(),
			})
	}

	// Get the locations
	locations, err := location.GetLocation(requestBody.DeviceIds, database)
	if err != nil {
		msg := "[Location] Failed to get location"
		log.Err(err).Msg(msg)
		return echoContext.JSON(http.StatusInternalServerError,
			map[string]any{
				"message": msg,
				"error":   err.Error(),
			})
	}

	return echoContext.JSON(http.StatusOK, locations)
}
