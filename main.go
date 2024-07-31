package main

import (
	"github.com/benmeehan/iot-metrics-service/config"
	"github.com/benmeehan/iot-metrics-service/handler"
	"github.com/benmeehan/iot-metrics-service/service"
	"github.com/benmeehan/iot-metrics-service/storage"

	"github.com/sirupsen/logrus"
)

func main() {
    // Initialize logger
    logger := logrus.New()
    logger.SetFormatter(&logrus.JSONFormatter{})

    // Load configuration
    cfg := config.LoadConfig()

    // Initialize InfluxDB client
    influxDBClient, err := storage.NewInfluxDBClient(cfg.InfluxDB, logger)
    if err != nil {
        logger.Fatalf("Failed to create InfluxDB client: %v", err)
    }
    defer influxDBClient.Close()

    // Initialize NATS client
    natsConn, err := service.NewNATSClient(cfg.NATS, logger)
    if err != nil {
        logger.Fatalf("Failed to connect to NATS: %v", err)
    }
    defer natsConn.Close()

    // Create a new handler
    metricsHandler := handler.NewMetricsHandler(influxDBClient, logger)

    // Subscribe to NATS topics
    err = service.SubscribeToMetrics(natsConn, metricsHandler, logger)
    if err != nil {
        logger.Fatalf("Failed to subscribe to NATS topics: %v", err)
    }

    // Keep the service running
    select {}
}
