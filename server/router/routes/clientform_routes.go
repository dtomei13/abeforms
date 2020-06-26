package routes

import (
	"github.com/austinlhx/server/controllers"
	"net/http"
)

var clientFormRoutes = []Route{
	Route{
		URI:          "/client/api/client",
		Method:       http.MethodPost,
		Handler:      controllers.AddClient,
		AuthRequired: false,
	},
}