package ioc

import "kafka-example/helpers"

var (
	application  string
	kafkaBrokers []string
	kafkaTopic   string
	kafkaGroup   string
)

func init() {
	application = "application"
	kafkaBrokers = helpers.GetEnvStringSlice("KAFKA_BROKERS", []string{"localhost:9092"})
	kafkaTopic = helpers.GetEnvString("KAFKA_TOPIC", "test")
	kafkaGroup = helpers.GetEnvString("KAFKA_GROUP", "test")
	singleton = Create()
}
