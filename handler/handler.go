package handler

import (
	"context"
	"encoding/json"

	"github.com/benmeehan/iot-metrics-service/models"
	"github.com/benmeehan/iot-metrics-service/storage"

	"github.com/sirupsen/logrus"
)

type MetricsHandler struct {
    dbClient storage.InfluxDBClient
    logger   *logrus.Logger
}

func NewMetricsHandler(dbClient storage.InfluxDBClient, logger *logrus.Logger) *MetricsHandler {
    return &MetricsHandler{dbClient: dbClient, logger: logger}
}

func (h *MetricsHandler) ProcessMetrics(data []byte) {
    // Parse metrics from data (assumed JSON for simplicity)
    var metrics models.Metrics
    if err := json.Unmarshal(data, &metrics); err != nil {
        h.logger.Errorf("Error unmarshalling metrics: %v", err)
        return
    }

    // Write metrics to InfluxDB
    err := h.dbClient.WriteMetrics(context.Background(), metrics)
    if err != nil {
        h.logger.Errorf("Error writing metrics to InfluxDB: %v", err)
    }
}
