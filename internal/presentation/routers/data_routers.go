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
}
