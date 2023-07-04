package main

import (
	"log"
	"net/http"

	"kafka-example/api"
	"kafka-example/helpers"
)

func main() {
	port := helpers.GetEnvString("PORT", "8080")

	log.Println("Starting server on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, api.GetRoutes()))
}
