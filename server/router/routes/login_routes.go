package routes

import (
	"net/http"

	"../../controllers"
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
	Route{
		URI:          "/clientdashboard/api/signin",
		Method:       http.MethodPost,
		Handler:      controllers.GetClient,
		AuthRequired: false,
	},
	Route{ // I think this is fixed..... BUT HAVE TO CHECK
		URI:          "/clientdashboard/api/signup",
		Method:       http.MethodPost,
		Handler:      controllers.AddClient, //This actually leads to the function for adding cases.... HAVE TO FIX THIS
		AuthRequired: false,
	},
}
