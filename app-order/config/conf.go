package config

// noinspection ALL
const (
	HTTP_ADDR = "HTTP_ADDRESS"

	NATS_CLUSTER = "NATS_CLUSTER"
	NATS_CLIENT  = "NATS_CLIENT"
	NATS_ADDR    = "NATS_ADDRESS"
)

var defaultConfig = map[string]string {
	//Transport Config
	HTTP_ADDR: ":8001",

	// Streaming Configuration
	NATS_ADDR:    "nats://localhost:4222",
	NATS_CLUSTER: "test-cluster",
	NATS_CLIENT:  "natsio-test-01-order",
}