package routers

import (
	"Gym-Service/internal/services"
	"strings"
)

var auth_routes = Routes{
	Route{
		"GetSession",
		strings.ToUpper("Post"),
		"/login",
		services.GetSession,
	},
	Route{
		"EndSession",
		strings.ToUpper("Post"),
		"/logout",
		services.EndSession,
	},
}
