package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	safe_area "server/database/safe-area"
)

func GetSafeAreaHandler(echoContext echo.Context, database *mongo.Client) error {
	// Get the query parameter
	queryParam := echoContext.QueryParam("safeAreaId")
	if len(queryParam) == 0 {
		msg := "[Safe-Area] Empty request parameter"
		log.Error().Msg(msg)
		return echoContext.JSON(http.StatusBadRequest, map[string]any{
			"error": msg,
		})
	}

	// Query the safe area
	safeAreaResult, err := safe_area.GetSafeArea("abcd", database)
	if err != nil {
		msg := "[Safe-Area] Failed to get the safe area"
		log.Error().Msg(msg)
		return echoContext.JSON(http.StatusInternalServerError, map[string]any{
			"message": msg,
			"error":   err.Error(),
		})
	}

	return echoContext.JSON(http.StatusOK, safeAreaResult)
}
