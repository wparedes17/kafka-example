package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func GetRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		handler := route.HandlerFunc

		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}

	return router
}
