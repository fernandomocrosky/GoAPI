package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Make return a new router instance
func Make() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
