package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Logger init with pretty format and timestamp enabled
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).With().Timestamp().Logger()
	// Set the logger to use globally
	log.Logger = logger

	// Load environment variables from .env file
	projectRootPath, err := os.Getwd()
	if err != nil {
		log.Fatal().Err(err).Msg("[Server] Failed to get project directory")
	}
	err = godotenv.Load(projectRootPath + "/.env")
	if err != nil {
		log.Fatal().Err(err).Msg("[Server] Error loading .env file")
	}

	// Init echo instance
	echoServer := echo.New()
	echoServer.GET("/", func(echoContext echo.Context) error {
		return echoContext.String(http.StatusOK, "Hello, World!")
	})
	echoServer.Logger.Fatal(echoServer.Start(":8080"))
}
