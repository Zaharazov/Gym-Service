package routers

import (
	"Gym-Service/internal/services"
	"strings"
)

var data_routes = Routes{

	Route{
		"GetData",
		strings.ToUpper("Get"),
		"/frontend/{_:.*}",
		services.GetData,
	},

	Route{
		"CreateEvent",
		strings.ToUpper("Post"),
		"/events",
		services.CreateEvent,
	},

	Route{
		"CreateCoach",
		strings.ToUpper("Post"),
		"/admin/create/coach",
		services.CreateCoach,
	},

	Route{
		"CreateAdmin",
		strings.ToUpper("Post"),
		"/admin/create/admin",
		services.CreateAdmin,
	},

	Route{
		"RestoreBackup",
		strings.ToUpper("Post"),
		"/admin/restore",
		services.RestoreBackup,
	},

	Route{
		"CreateBackup",
		strings.ToUpper("Post"),
		"/admin/backup",
		services.CreateBackup,
	},
}
