package routes

import (
	"github.com/austinlhx/server/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	URI          string
	Method       string
	Handler      func(w http.ResponseWriter, r *http.Request)
	AuthRequired bool
}

// Load the routes
func Load() []Route {
	routes := dashboardRoutes
	routes = append(routes, clientFormRoutes...)
	routes = append(routes, loginRoutes...)
	return routes
}

func SetupRoutesWithMiddlewares(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		if route.AuthRequired {
			r.HandleFunc(route.URI,
				middlewares.SetMiddlewareLogger(
					middlewares.SetMiddlewareJSON(
						middlewares.SetMiddlewareAuthentication(route.Handler),
					),
				),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI,
				middlewares.SetMiddlewareLogger(
					middlewares.SetMiddlewareJSON(route.Handler),
				),
			).Methods(route.Method)
		}

	}
	return r
}
