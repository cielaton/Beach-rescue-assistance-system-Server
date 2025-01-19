package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	"server/database/user"
)

func GetUserHandler(echoContext echo.Context, database *mongo.Client) error {
	// Get the query parameter
	queryParam := echoContext.QueryParam("userName")
	if len(queryParam) == 0 {
		msg := "[User] Empty request parameter"
		log.Error().Msg(msg)
		return echoContext.JSON(http.StatusBadRequest, map[string]any{
			"error": msg,
		})
	}

	userResult, err := user.GetUser(queryParam, database)
	if err != nil {
		msg := "[User] Failed to get user info"
		log.Error().Msg(msg)
		return echoContext.JSON(http.StatusInternalServerError, map[string]any{
			"message": msg,
			"error":   err.Error(),
		})
	}

	return echoContext.JSON(http.StatusOK, userResult)
}
