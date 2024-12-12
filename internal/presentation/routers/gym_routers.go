package routers

import (
	"Gym-Service/internal/services"
	"strings"
)

var gyms_routes = Routes{

	Route{
		"SetGym",
		strings.ToUpper("Post"),
		"/set_gym",
		services.SetGym,
	},
}
