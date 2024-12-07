package routers

import (
	"Gym-Service/internal/services"
	"strings"
)

var register_routes = Routes{
	Route{
		"RegUser",
		strings.ToUpper("Post"),
		"/register",
		services.RegUser,
	},
}
