package handler

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	"server/database/device"
)

func GetDeviceHandler(echoContext echo.Context, database *mongo.Client) error {
	// Get the query parameter
	queryParam := echoContext.QueryParam("deviceId")
	if len(queryParam) == 0 {
		return echoContext.JSON(http.StatusBadRequest, map[string]any{
			"message": "Empty query parameter",
		})
	}
	deviceResult, err := device.GetDevice("abcd", database)

	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, err.Error())
	}

	return echoContext.JSON(http.StatusOK, deviceResult)
}
