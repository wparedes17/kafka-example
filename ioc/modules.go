package ioc

import "kafka-example/helpers"

var (
	application  string
	kafkaBrokers []string
	kafkaTopic   string
)

func init() {
	application = "application"
	kafkaBrokers = helpers.GetEnvStringSlice("KAFKA_BROKERS", []string{"localhost:9092"})
	kafkaTopic = helpers.GetEnvString("KAFKA_TOPIC", "test")

	singleton = Create()
}
