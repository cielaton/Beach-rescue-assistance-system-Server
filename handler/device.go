package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	"server/database/device"
	"server/model"
)

func GetDeviceByIdHandler(echoContext echo.Context, database *mongo.Client) error {
	// Get the path parameter
	pathParam := echoContext.Param("deviceId")
	if len(pathParam) == 0 {
		msg := "[Device] Empty path parameter"
		log.Error().Msg(msg)
		return echoContext.JSON(http.StatusBadRequest, map[string]any{
			"error": msg,
		})
	}

	deviceResult, err := device.GetDeviceByDeviceId(pathParam, database)

	if err != nil {
		msg := "[Device] Failed to get device info"
		return echoContext.JSON(http.StatusInternalServerError, map[string]any{
			"message": msg,
			"error":   err.Error(),
		})
	}

	return echoContext.JSON(http.StatusOK, deviceResult)
}

func GetDeviceBySafeAreaIdHandler(echoContext echo.Context, database *mongo.Client) error {
	// Get the path parameter
	pathParam := echoContext.Param("safeAreaId")
	if len(pathParam) == 0 {
		msg := "[Device] Empty path parameter"
		log.Error().Msg(msg)
		return echoContext.JSON(http.StatusBadRequest, map[string]any{
			"error": msg,
		})
	}

	deviceResult, err := device.GetDeviceBySafeAreaId(pathParam, database)

	if err != nil {
		msg := "[Device] Failed to get device info"
		return echoContext.JSON(http.StatusInternalServerError, map[string]any{
			"message": msg,
			"error":   err.Error(),
		})
	}

	return echoContext.JSON(http.StatusOK, deviceResult)
}

func DeleteDeviceHandler(echoContext echo.Context, database *mongo.Client) error {
	// Get the path parameter
	pathParam := echoContext.Param("deviceId")
	if len(pathParam) == 0 {
		msg := "[Device] Empty path parameter"
		log.Error().Msg(msg)
		return echoContext.JSON(http.StatusBadRequest, map[string]any{
			"error": msg,
		})
	}

	err := device.DeleteDevice(pathParam, database)
	if err != nil {
		msg := "[Device] Failed to delete device"
		log.Error().Msg(msg)
		return echoContext.JSON(http.StatusInternalServerError, map[string]any{
			"message": msg,
			"error":   err.Error(),
		})
	}
	return echoContext.JSON(http.StatusOK, map[string]any{
		"message": "Device deleted",
	})
}

func ChangeDeviceActiveStatus(echoContext echo.Context, database *mongo.Client) error {
	// Bind the POST request body
	var requestBody model.DeviceActiveChangeRequest
	err := echoContext.Bind(&requestBody)
	if err != nil {
		msg := "[Device] Invalid request body"
		return echoContext.JSON(http.StatusBadRequest,
			map[string]any{
				"message": msg,
				"error":   err.Error(),
			})
	}

	err = device.ChangeDeviceActiveStatus(requestBody.DeviceId, requestBody.Status, database)
	if err != nil {
		msg := "[Device] Failed to change device status"
		return echoContext.JSON(http.StatusInternalServerError, map[string]any{
			"message": msg,
			"error":   err.Error(),
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]any{
		"message": "Device status changed",
	})
}
