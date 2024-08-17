package rotas

import (
	"api-devbook/src/controllers"
	"net/http"
)

var rotaUsuarios = []Rota{
	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CreateUser,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.FindUsers,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{userId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.FindUser,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{userId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.UpdateUser,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{userId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeleteUser,
		RequerAutenticacao: false,
	},
}
