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

	// API routes
	// Root
	echoServer.GET("/", func(echoContext echo.Context) error {
		return echoContext.String(http.StatusOK, "Hello, World!")
	})
	// Device
	echoServer.GET("/device/byDeviceId/:deviceId", func(echoContext echo.Context) error {
		return handler.GetDeviceByIdHandler(echoContext, databaseClient)
	})
	echoServer.GET("/device/bySafeAreaId/:safeAreaId", func(echoContext echo.Context) error {
		return handler.GetDeviceBySafeAreaIdHandler(echoContext, databaseClient)
	})
	echoServer.DELETE("/device/:deviceId", func(echoContext echo.Context) error {
		return handler.DeleteDeviceHandler(echoContext, databaseClient)
	})
	echoServer.POST("/device/activeStatus", func(echoContext echo.Context) error {
		return handler.ChangeDeviceActiveStatus(echoContext, databaseClient)
	})
	// User
	echoServer.GET("/user", func(echoContext echo.Context) error {
		return handler.GetUserHandler(echoContext, databaseClient)
	})
	// Location
	echoServer.POST("/location", func(echoContext echo.Context) error {
		return handler.GetLocationHandler(echoContext, databaseClient)
	})
	// Safe Area
	echoServer.GET("/safe-area", func(echoContext echo.Context) error {
		return handler.GetSafeAreaHandler(echoContext, databaseClient)
	})
	// Rescuer
	echoServer.GET("/rescuer/byId/:rescuerId", func(echoContext echo.Context) error {
		return handler.GetRescuerByRescuerIdHandler(echoContext, databaseClient)
	})
	echoServer.GET("/rescuer/bySafeAreaId/:safeAreaId", func(echoContext echo.Context) error {
		return handler.GetRescuerBySafeAreaIdHandler(echoContext, databaseClient)
	})
	// Rescuer Location
	echoServer.POST("/rescuer-location", func(echoContext echo.Context) error {
		return handler.GetRescuerLocationHandler(echoContext, databaseClient)
	})

	// Start the server
	echoServer.Logger.Fatal(echoServer.Start(":8080"))
	// Disconnect the database
	if databaseClient != nil {
		err = databaseClient.Disconnect(context.Background())
		if err != nil {
			panic(err)
		}
	}
}
