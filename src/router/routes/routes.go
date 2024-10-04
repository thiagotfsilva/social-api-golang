package routes

import (
	"api-devbook/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI            string
	Method         string
	Function       func(http.ResponseWriter, *http.Request)
	Authentication bool
}

// Config coloca todas as rotas dentro do router
func Config(r *mux.Router) *mux.Router {
	routes := UserRoute
	routes = append(routes, LoginRoute)
	routes = append(routes, PublicationRoute...)

	for _, route := range routes {
		if route.Authentication {
			r.HandleFunc(
				route.URI,
				middlewares.Logger(middlewares.HandleAuth(route.Function)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(
				route.URI,
				middlewares.Logger(route.Function),
			).Methods(route.Method)
		}
	}

	return r
}
