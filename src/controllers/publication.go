package controllers

import (
	"api-devbook/src/infra"
	"api-devbook/src/models"
	"api-devbook/src/repositories"
	"api-devbook/src/utils/auth"
	"api-devbook/src/utils/response"
	"encoding/json"
	"io"
	"net/http"
)

func CreatePublication(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.ExtractUserId(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publication models.Publication
	if err = json.Unmarshal(bodyRequest, &publication); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	publication.AuthorId = userId
	if err = publication.Prepare(); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := infra.Connect()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	publicationRepository := repositories.NewPublicationRepository(db)
	_, err = publicationRepository.Create(publication)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, publication)
}
func FindPublications(w http.ResponseWriter, r *http.Request) {

}
func FindPublication(w http.ResponseWriter, r *http.Request) {

}
func UpdatePublication(w http.ResponseWriter, r *http.Request) {

}
func DeletePublication(w http.ResponseWriter, r *http.Request) {

}
