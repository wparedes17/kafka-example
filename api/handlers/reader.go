package handlers

import (
	"context"
	"kafka-example/api/utils"
	"kafka-example/ioc"
	"net/http"
)

func ReaderHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	kafkaService := ioc.New().GetKafkaService()
	msg, err := kafkaService.Read(ctx)
	if err != nil {
		panic(err)
	}

	utils.WriteResponse(ctx, w, http.StatusOK, map[string]string{
		"key":     "You have a new message",
		"message": msg,
	})
}
