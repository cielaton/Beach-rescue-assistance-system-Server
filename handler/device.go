package handler

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	"server/database/device"
)

func GetDeviceHandler(echoContext echo.Context, database *mongo.Client) error {
	deviceResult, err := device.GetDevice(database)

	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, err.Error())
	}

	return echoContext.JSON(http.StatusOK, deviceResult)
}
