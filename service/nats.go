package service

import (
	"github.com/benmeehan/iot-metrics-service/handler"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

func NewNATSClient(cfg struct {
    URL string
}, logger *logrus.Logger) (*nats.Conn, error) {
    opts := []nats.Option{nats.Name("NATS Client")}
    
    // Optionally add TLS configuration here
    // opts = append(opts, nats.Secure(tlsConfig))

    conn, err := nats.Connect(cfg.URL, opts...)
    if err != nil {
        logger.Errorf("Error connecting to NATS: %v", err)
        return nil, err
    }
    return conn, nil
}

func SubscribeToMetrics(nc *nats.Conn, handler *handler.MetricsHandler, logger *logrus.Logger) error {
    _, err := nc.Subscribe("metrics", func(msg *nats.Msg) {
        handler.ProcessMetrics(msg.Data)
    })
    if err != nil {
        logger.Errorf("Error subscribing to NATS topic: %v", err)
    }
    return err
}
