package routes

import (
	"github.com/austinlhx/server/controllers"
	"net/http"
)

var loginRoutes = []Route{
	Route{
		URI:          "/lawyerdashboard/api/signin",
		Method:       http.MethodPost,
		Handler:      controllers.GetLawyer,
		AuthRequired: false,
	},
	Route{
		URI:          "/lawyerdashboard/api/signup",
		Method:       http.MethodPost,
		Handler:      controllers.AddLawyer,
		AuthRequired: false,
	},
}

