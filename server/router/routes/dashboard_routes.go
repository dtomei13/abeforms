package routes

import (
	"github.com/austinlhx/server/controllers"
	"net/http"
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
		URI: "/lawyerdashboard/api/takecase/{id}",
		Method: http.MethodPut,
		Handler: controllers.TakeCase,
		AuthRequired: true,
	},
	Route{
		URI: "/lawyerdashboard/api/schedulemeeting",
		Method: http.MethodPost,
		Handler: controllers.CreateMeeting,
		AuthRequired: true,
	},

}
