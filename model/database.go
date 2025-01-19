package model

import "time"

type User struct {
	UserId      string `json:"userId"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	WorkingSite string `json:"workingSite"`
}

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

type SafeAreaLocation struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
type SafeArea struct {
	SafeAreaId   string             `json:"safeAreaId"`
	DeviceIds    []string           `json:"deviceIds"`
	LocationName string             `json:"locationName"`
	Boundary     []SafeAreaLocation `json:"boundary"`
}
