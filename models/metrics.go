package models

type Metrics struct {
	DeviceID   string  `json:"device_id"`
	Timestamp  string  `json:"timestamp"`
	MetricType string  `json:"metric_type"`
	Value      float64 `json:"value"`
}
