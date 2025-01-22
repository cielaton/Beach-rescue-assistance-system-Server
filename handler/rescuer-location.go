package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	"server/database/rescuer-location"
	"server/model"
)

func GetRescuerLocationHandler(echoContext echo.Context, database *mongo.Client) error {
	// Bind the POST request body
	var requestBody model.RescuerLocationRequest
	err := echoContext.Bind(&requestBody)
	if err != nil {
		msg := "[Rescuer Location] Invalid request body"
		log.Err(err).Msg(msg)
		return echoContext.JSON(http.StatusBadRequest,
			map[string]any{
				"message": msg,
				"error":   err.Error(),
			})
	}

	// Get the locations
	rescuerLocations, err := rescuer_location.GetRescuerLocation(requestBody.RescuerIds, database)
	if err != nil {
		msg := "[Rescuer Location] Failed to get rescuer location"
		log.Err(err).Msg(msg)
		return echoContext.JSON(http.StatusInternalServerError,
			map[string]any{
				"message": msg,
				"error":   err.Error(),
			})
	}

	return echoContext.JSON(http.StatusOK, rescuerLocations)
}
