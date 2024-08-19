package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	RequiredAuth bool
}

// Config coloca todas as rotas dentro do router
func Config(r *mux.Router) *mux.Router {
	routes := UserRoute
	routes = append(routes, LoginRoute)

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
