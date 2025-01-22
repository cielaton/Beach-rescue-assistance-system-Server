package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	"server/database/rescuer"
)

func GetRescuerByRescuerIdHandler(echoContext echo.Context, database *mongo.Client) error {
	// Get the query parameter
	queryParam := echoContext.Param("rescuerId")
	if len(queryParam) == 0 {
		msg := "[Device] Empty request parameter"
		log.Error().Msg(msg)
		return echoContext.JSON(http.StatusBadRequest, map[string]any{
			"error": msg,
		})
	}

	rescuerResult, err := rescuer.GetRescuerByRescuerId(queryParam, database)

	if err != nil {
		msg := "[Device] Failed to get device info"
		log.Error().Msg(msg)
		return echoContext.JSON(http.StatusInternalServerError, map[string]any{
			"message": msg,
			"error":   err.Error(),
		})
	}

	return echoContext.JSON(http.StatusOK, rescuerResult)
}

func GetRescuerBySafeAreaIdHandler(echoContext echo.Context, database *mongo.Client) error {
	// Get the query parameter
	queryParam := echoContext.Param("safeAreaId")
	fmt.Println(queryParam)
	if len(queryParam) == 0 {
		msg := "[Device] Empty request parameter"
		log.Error().Msg(msg)
		return echoContext.JSON(http.StatusBadRequest, map[string]any{
			"error": msg,
		})
	}

	rescuerResult, err := rescuer.GetRescuerBySafeAreaId(queryParam, database)

	if err != nil {
		msg := "[Rescuer] Failed to get rescuer info"
		log.Error().Msg(msg)
		return echoContext.JSON(http.StatusInternalServerError, map[string]any{
			"message": msg,
			"error":   err.Error(),
		})
	}

	return echoContext.JSON(http.StatusOK, rescuerResult)
}
