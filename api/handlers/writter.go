package handlers

import (
	"kafka-example/api/utils"
	"kafka-example/ioc"
	"kafka-example/models"
	"net/http"
)

func WriterHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	msgRequest := models.MsgRequest{}
	err := utils.DecodeSchema(&msgRequest, r.URL.Query())
	if err != nil {
		panic(err)
	}

	msgText, err := utils.GenerateJSONString(&msgRequest)
	if err != nil {
		panic(err)
	}

	kafkaService := ioc.New().GetKafkaService()
	err = kafkaService.Write(ctx, msgRequest.MsgKey, msgText)
	if err != nil {
		panic(err)
	}

	utils.WriteResponse(ctx, w, http.StatusOK, map[string]string{
		"key":    msgRequest.MsgKey,
		"status": "Message sent successfully",
	})
}
