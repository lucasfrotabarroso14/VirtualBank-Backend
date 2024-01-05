package router

import (
	"github.com/gorilla/mux"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/router/routes"
)

func Gerar() *mux.Router {
	r := mux.NewRouter()
	return routes.ConfigRoutes(r)
}
