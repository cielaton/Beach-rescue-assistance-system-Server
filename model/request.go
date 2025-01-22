package model

type LocationRequest struct {
	DeviceIds []string `json:"deviceIds"`
}

type RescuerLocationRequest struct {
	RescuerIds []string `json:"rescuerIds"`
}
