package handlers

import (
	"kafka-example/api/utils"
	"net/http"
)

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	utils.WriteResponse(ctx, w, http.StatusOK, map[string]string{
		"key":     "Health check",
		"message": "API is up and running",
	})
}
