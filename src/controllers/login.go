package controllers

import (
	"api-devbook/src/infra"
	"api-devbook/src/models"
	"api-devbook/src/repositories"
	"api-devbook/src/utils/auth"
	handlehash "api-devbook/src/utils/handleHash"
	"api-devbook/src/utils/response"
	"encoding/json"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	db, err := infra.Connect()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)

	userExist, err := userRepository.FindByEmail(user.Email)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	err = handlehash.VerifyPassword(userExist.Password, user.Password)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.CreateToken(userExist.ID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	w.Write([]byte(token))
}
