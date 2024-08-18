package routes

import (
	"api-devbook/src/controllers"
	"net/http"
)

var UserRoute = []Route{
	{
		URI:          "/usuarios",
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		RequiredAuth: false,
	},
	{
		URI:          "/usuarios",
		Method:       http.MethodGet,
		Function:     controllers.FindUsers,
		RequiredAuth: false,
	},
	{
		URI:          "/usuarios/{userId}",
		Method:       http.MethodGet,
		Function:     controllers.FindUser,
		RequiredAuth: false,
	},
	{
		URI:          "/usuarios/{userId}",
		Method:       http.MethodPut,
		Function:     controllers.UpdateUser,
		RequiredAuth: false,
	},
	{
		URI:          "/usuarios/{userId}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteUser,
		RequiredAuth: false,
	},
}
