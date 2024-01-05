package routes

import (
	"github.com/gorilla/mux"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/middlewares"
	"net/http"
)

type Route struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	NeedAuth bool
}

func ConfigRoutes(r *mux.Router) *mux.Router {
	var routes = AccountRoutes
	routes = append(routes, loginRoute)

	for _, route := range routes {
		r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)

	}
	return r
}
