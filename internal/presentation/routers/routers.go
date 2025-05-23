package routers

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

type Routes []Route

var routes Routes

func NewRouter() *mux.Router {

	routes = append(routes, gyms_routes...)
	routes = append(routes, passes_routes...)
	routes = append(routes, data_routes...)
	routes = append(routes, pages_routes...)
	routes = append(routes, auth_routes...)
	routes = append(routes, register_routes...)

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		//handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
