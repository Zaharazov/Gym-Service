package routers

import (
	"Gym-Service/internal/services"
	"strings"
)

var passes_routes = Routes{

	Route{
		"SetPass",
		strings.ToUpper("Post"),
		"/set_pass",
		services.SetPass,
	},
}
