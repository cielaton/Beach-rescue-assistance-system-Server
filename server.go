package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"server/database"
	"server/handler"
	"time"
)

func main() {
	// Init echo instance
	echoServer := echo.New()
	// Make the server recover on panic
	echoServer.Use(middleware.Recover())

	// Logger init with pretty format and timestamp enabled
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).With().Timestamp().Logger()
	// Set the logger to use globally
	log.Logger = logger
	// Logger configuration which enable necessary fields
	echoServer.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogLatency:  true,
		LogProtocol: true,
		LogURI:      true,
		LogStatus:   true,
		LogRemoteIP: true,
		LogValuesFunc: func(echoContext echo.Context, loggerValues middleware.RequestLoggerValues) error {
			logger.Info().Str("Protocol", loggerValues.Protocol).Str("Address", loggerValues.RemoteIP).Str("URI", loggerValues.URI).Int("Status", loggerValues.Status).Dur("Latency", loggerValues.Latency).Msg("Request")
			return nil
		},
	}))

	// Load environment variables from .env file
	projectRootPath, err := os.Getwd()
	if err != nil {
		log.Fatal().Err(err).Msg("[Server] Failed to get project directory")
	}
	err = godotenv.Load(projectRootPath + "/.env")
	if err != nil {
		log.Fatal().Err(err).Msg("[Server] Error loading .env file")
	}
	log.Info().Msg("[Server] Successfully loaded .env file")

	// Database connect
	databaseClient, err := database.Connect()
	if err != nil {
		log.Fatal().Msg("[Server] Error connecting to database")
	}

	echoServer.GET("/", func(echoContext echo.Context) error {
		return echoContext.String(http.StatusOK, "Hello, World!")
	})
	echoServer.GET("/device", func(echoContext echo.Context) error {
		return handler.GetDeviceHandler(echoContext, databaseClient)
	})
	echoServer.Logger.Fatal(echoServer.Start(":8080"))

	// Disconnect the database
	if databaseClient != nil {
		err = databaseClient.Disconnect(context.Background())
		if err != nil {
			panic(err)
		}
	}
}
