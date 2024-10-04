package routes

import (
	"api-devbook/src/controllers"
	"net/http"
)

var PublicationRoute = []Route{
	{
		URI:            "/publications",
		Method:         http.MethodPost,
		Function:       controllers.CreatePublication,
		Authentication: true,
	},
	{
		URI:            "/publications",
		Method:         http.MethodGet,
		Function:       controllers.FindPublications,
		Authentication: true,
	},
	{
		URI:            "/publications/{publicationId}",
		Method:         http.MethodGet,
		Function:       controllers.FindPublication,
		Authentication: true,
	},
	{
		URI:            "/publications/{publicationId}",
		Method:         http.MethodPut,
		Function:       controllers.UpdatePublication,
		Authentication: true,
	},
	{
		URI:            "/publications/{publicationId}",
		Method:         http.MethodDelete,
		Function:       controllers.DeletePublication,
		Authentication: true,
	},
}
