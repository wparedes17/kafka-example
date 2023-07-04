package api

import (
	"kafka-example/api/handlers"
)

var routes = []Route{
	{
		"Healthcheck Endpoint",
		"GET",
		"/healthcheck",
		handlers.Healthcheck,
	},
	{
		"Writer Endpoint",
		"GET",
		"/writer",
		handlers.WriterHandler,
	},
	{
		"Reader Endpoint",
		"GET",
		"/reader",
		handlers.ReaderHandler,
	},
}
