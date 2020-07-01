package routes

import (
	"net/http"

	"../../controllers"
)

var dashboardRoutes = []Route{
	Route{
		URI:          "/lawyerdashboard/api/mycases",
		Method:       http.MethodGet,
		Handler:      controllers.GetEachCase,
		AuthRequired: true,
	},
	Route{
		URI:          "/lawyerdashboard/api/opencases",
		Method:       http.MethodGet,
		Handler:      controllers.GetCase,
		AuthRequired: true,
	},
	Route{
		URI:          "/lawyerdashboard/api/takecase/{id}",
		Method:       http.MethodPost,
		Handler:      controllers.TakeCase,
		AuthRequired: true,
	},
	Route{
		URI:          "/lawyerdashboard/api/schedulemeeting",
		Method:       http.MethodPost,
		Handler:      controllers.CreateMeeting,
		AuthRequired: true,
	},
	Route{
		URI:          "/clientdashboard/api/unassignedcases",
		Method:       http.MethodGet,
		Handler:      controllers.GetUnassignedCase,
		AuthRequired: true,
	},
	Route{
		URI:          "/clientdashboard/api/assignedcases",
		Method:       http.MethodGet,
		Handler:      controllers.GetAssignedCase,
		AuthRequired: true,
	},
}
