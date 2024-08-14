package controllers

import (
	"api-devbook/src/infra"
	"api-devbook/src/models"
	"api-devbook/src/repositories"
	"api-devbook/src/utils/response"
	"encoding/json"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Prepare(); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := infra.Connect()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	user.ID, err = userRepository.Create(user)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
	}

	response.JSON(w, http.StatusCreated, user)
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Usuarios"))
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Usuario"))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizar Usuario"))
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletar Usuario"))
}
