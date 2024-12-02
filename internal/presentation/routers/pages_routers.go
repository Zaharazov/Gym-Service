package routers

import (
	"Gym-Service/internal/services"
	"strings"
)

var pages_routes = Routes{
	Route{
		"GetAuthPage",
		strings.ToUpper("Get"),
		"/",
		services.GetAuthPage,
	},

	Route{
		"GetGymsPage",
		strings.ToUpper("Get"),
		"/gyms",
		services.GetGymsPage,
	},

	Route{
		"GetGymPageById",
		strings.ToUpper("Get"),
		"/gyms/{gym_id}",
		services.GetGymPageById,
	},
}
