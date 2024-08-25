package routes

import (
	"api-devbook/src/controllers"
	"net/http"
)

var LoginRoute = Route{
	URI:            "/login",
	Method:         http.MethodPost,
	Function:       controllers.Login,
	Authentication: false,
}
