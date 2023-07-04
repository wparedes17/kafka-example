package utils

import (
	"context"
	"encoding/json"
	"net/http"
)

func WriteResponse(ctx context.Context, rw http.ResponseWriter, statusCode int, body interface{}) error {
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.WriteHeader(statusCode)

	err := json.NewEncoder(rw).Encode(body)
	if err != nil {
		panic(err)
	}

	return err
}
