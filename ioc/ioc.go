package ioc

import (
	"log"
	"net/http"

	"kafka-example/services/kafkaservice"
)

type Container interface {
	GetApplication() string
	GetKafkaService() kafkaservice.KafkaService
}

var singleton container

func New() Container {
	return &singleton
}

type container struct {
	httpClient   *http.Client
	kafkaService kafkaservice.KafkaService
}

func Create() container {
	log.Println("Initializing application singletons...")
	c := container{}

	log.Println("Creating HTTP Client")
	c.httpClient = http.DefaultClient

	log.Println("Initializing Kafka Service")
	c.kafkaService = kafkaservice.New(kafkaTopic, kafkaGroup, kafkaBrokers...)

	log.Println("Done initializing singletons!")
	return c
}

func (c *container) GetHttpClient() *http.Client {
	return c.httpClient
}

func (c *container) GetApplication() string {
	return application
}

func (c *container) GetKafkaService() kafkaservice.KafkaService {
	return c.kafkaService
}
