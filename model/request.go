package model

type LocationRequest struct {
	DeviceIds []string `json:"deviceIds"`
}

type RescuerLocationRequest struct {
	RescuerIds []string `json:"rescuerIds"`
}

type DeviceActiveChangeRequest struct {
	DeviceId string `json:"deviceId"`
	Status   bool   `json:"status"`
}
