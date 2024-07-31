# IoT Metrics Service

## Overview

The IoT Metrics Service is a microservice designed to receive heartbeat and metrics data from IoT devices, process it, and store it in InfluxDB for future analytics. The service is built using Go and utilizes InfluxDB for time-series data storage. It also integrates with NATS or Kafka for messaging and has support for logging and security.

## Features

- **Receive Metrics:** Collect heartbeat and metrics data from IoT devices.
- **Time-Series Storage:** Store metrics in InfluxDB for efficient querying and analysis.
- **Messaging Integration:** Supports messaging via NATS or Kafka.
- **Logging:** Integrated logging for debugging and monitoring.
- **Security:** Basic security practices for secure communication.

## Architecture

- **Messaging Queue:** NATS or Kafka (configurable)
- **Storage:** InfluxDB for time-series data
- **Logging:** Logrus for structured logging
- **Language:** Go

## Setup

### Prerequisites

- Go 1.18 or higher
- InfluxDB instance
- NATS or Kafka instance (optional, for messaging)
- Go modules

### Installation

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/your-username/iot-metrics-service.git
   cd iot-metrics-service
   ```

2. **Install Dependencies:**

   ```bash
   go mod tidy
   ```

3. **Configure InfluxDB:**

   Create a configuration file (e.g., `config.yaml`) with the following structure:

   ```yaml
   url: "http://localhost:8086"
   token: "your-token"
   org: "your-org"
   bucket: "your-bucket"
   ```

   Ensure that your InfluxDB instance is running and accessible.

4. **Set Up Environment Variables (if applicable):**

   You may need to set environment variables for your configuration or sensitive data.

5. **Build the Project:**

   ```bash
   go build -o iot-metrics-service
   ```

6. **Run the Service:**

   ```bash
   ./iot-metrics-service
   ```

## Usage

1. **Configure the Messaging Queue:**

   If using NATS or Kafka, configure the service to connect to your messaging queue by setting the appropriate environment variables or updating the configuration file.

2. **Send Metrics:**

   Metrics can be sent to the service via your configured messaging queue. The service will process and store these metrics in InfluxDB.

3. **Query Metrics:**

   Use InfluxDB's query interface or CLI to query and analyze the stored metrics.

## Example Configuration

Here's an example configuration file (`config.yaml`):

```yaml
url: "http://localhost:8086"
token: "my-secret-token"
org: "my-org"
bucket: "my-bucket"
```

## Code Structure

- `main.go`: Entry point for the service.
- `storage/influxdb.go`: Implementation for interacting with InfluxDB.
- `models/models.go`: Data models for metrics.
- `config/config.yaml`: Configuration file (template).

## Logging

The service uses Logrus for logging. Logs are output to the console by default. You can customize logging behavior by configuring Logrus settings.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For questions or support, please contact [benmeehan111@example.com](mailto:benmeehan111@example.com).
