package model

type Gyroscope struct {
	Device
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}
