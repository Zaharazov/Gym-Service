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

	Route{
		"GetGymsPassesPage",
		strings.ToUpper("Get"),
		"/gyms_passes",
		services.GetGymsPassesPage,
	},

	Route{
		"GetEqipmentPage",
		strings.ToUpper("Get"),
		"/equipment",
		services.GetEquipmentPage,
	},

	Route{
		"GetAdminPage",
		strings.ToUpper("Get"),
		"/admin",
		services.GetAdminPage,
	},

	Route{
		"GetAdminPage",
		strings.ToUpper("Post"),
		"/admin",
		services.GetAdminPage,
	},

	Route{
		"NotFoundPage",
		strings.ToUpper("GET"),
		"/{_:.*}",
		services.NotFoundPage,
	},
}
