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
