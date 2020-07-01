package router

import (
	"./routes"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutesWithMiddlewares(router)

}
