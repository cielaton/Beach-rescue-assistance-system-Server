package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	"server/database/user"
)

func GetUserHandler(echoContext echo.Context, database *mongo.Client) error {
	// Get the query parameter
	queryParam := echoContext.QueryParam("userName")
	fmt.Printf("queryParam: %s\n", queryParam)
	if len(queryParam) == 0 {
		return echoContext.JSON(http.StatusBadRequest, map[string]any{
			"message": "Empty query parameter",
		})
	}

	userResult, err := user.GetUser(queryParam, database)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, err.Error())
	}

	return echoContext.JSON(http.StatusOK, userResult)
}
