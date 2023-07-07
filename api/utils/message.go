package utils

import (
	"encoding/json"
	"kafka-example/models"
)

func GenerateJSONString(msgRequest *models.MsgRequest) (string, error) {
	message := models.MessageToProcess{
		RetryNumber:  0,
		JobID:        msgRequest.JobID,
		UserID:       msgRequest.UserID,
		ResumeID:     msgRequest.ResumeID,
		BoardID:      nil,
		Origin:       0,
		Rank:         0,
		CreationDate: "",
		Skills:       []models.UserSkills{},
	}

	jsonString, err := json.Marshal(message)
	if err != nil {
		return "", err
	}

	return string(jsonString), nil
}
