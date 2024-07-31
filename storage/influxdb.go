package storage

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/benmeehan/iot-metrics-service/models"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/sirupsen/logrus"
)

type InfluxDBClient interface {
    WriteMetrics(ctx context.Context, metrics models.Metrics) error
    Close() error
}

type influxDBClient struct {
    client   influxdb2.Client
    writeAPI api.WriteAPIBlocking
    logger   *logrus.Logger
}

func NewInfluxDBClient(cfg struct {
    URL    string
    Token  string
    Org    string
    Bucket string
}, logger *logrus.Logger) (InfluxDBClient, error) {
    // Ensure URL uses HTTPS
    if !strings.HasPrefix(cfg.URL, "https://") {
        return nil, fmt.Errorf("URL must use HTTPS: %s", cfg.URL)
    }

    client := influxdb2.NewClient(cfg.URL, cfg.Token)
    writeAPI := client.WriteAPIBlocking(cfg.Org, cfg.Bucket)
    return &influxDBClient{
        client:   client,
        writeAPI: writeAPI,
        logger:   logger,
    }, nil
}

func (c *influxDBClient) WriteMetrics(ctx context.Context, metrics models.Metrics) error {
    p := influxdb2.NewPointWithMeasurement(metrics.MetricType).
        AddTag("device_id", metrics.DeviceID).
        AddField("value", metrics.Value).
        SetTime(parseTime(metrics.Timestamp))

    // Write point to InfluxDB
    if err := c.writeAPI.WritePoint(ctx, p); err != nil {
        c.logger.Errorf("Error writing point to InfluxDB: %v", err)
        return err
    }
    return nil
}

func parseTime(timestamp string) time.Time {
    t, err := time.Parse(time.RFC3339, timestamp)
    if err != nil {
        // Handle parsing error (log or return a default value)
        return time.Now() // Default to current time if parsing fails
    }
    return t
}

func (c *influxDBClient) Close() error {
    c.client.Close()
    return nil
}
