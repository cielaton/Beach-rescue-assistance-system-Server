package model

import "time"

type Location struct {
	DeviceId      string    `json:"deviceId"`
	DatePublished time.Time `json:"datePublished"`
	Description   string    `json:"description"`
	Longitude     float64   `json:"longitude"`
	Latitude      float64   `json:"latitude"`
	StatusCode    int       `json:"statusCode"`
}

type Device struct {
	DeviceId   string    `json:"deviceId"`
	SafeAreaId string    `json:"datePublished"`
	DateAdded  time.Time `json:"dateAdded"`
	PrivateKey string    `json:"privateKey"`
	IsEnabled  bool      `json:"isEnabled"`
}
