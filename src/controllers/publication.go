package controllers

import (
	"api-devbook/src/infra"
	"api-devbook/src/models"
	"api-devbook/src/repositories"
	"api-devbook/src/utils/auth"
	"api-devbook/src/utils/response"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	userId, err := auth.ExtractUserId(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	db, err := infra.Connect()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	publicationRepository := repositories.NewPublicationRepository(db)
	publications, err := publicationRepository.Fetch(userId)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, publications)
}

func FindPublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationId, err := strconv.ParseUint(params["publicationId"], 10, 64)
	if err != nil {
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
	publication, err := publicationRepository.FindById(publicationId)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, publication)
}

func UpdatePublication(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.ExtractUserId(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	publicationId, err := strconv.ParseUint(params["publicationId"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := infra.Connect()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
	}
	defer db.Close()

	publicationRepository := repositories.NewPublicationRepository(db)
	publicatioExist, err := publicationRepository.FindById(publicationId)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if publicatioExist.AuthorId != userId {
		response.Erro(w, http.StatusForbidden, errors.New("Forbiden"))
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

	if err = publication.Prepare(); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = publicationRepository.Update(publicationId, publication); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)

}

func DeletePublication(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.ExtractUserId(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	publicationId, err := strconv.ParseUint(params["publicationId"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := infra.Connect()
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	defer db.Close()

	publicationRepository := repositories.NewPublicationRepository(db)
	publicationExist, err := publicationRepository.FindById(publicationId)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if publicationExist.AuthorId != userId {
		response.Erro(w, http.StatusForbidden, errors.New("Forbiden"))
		return
	}

	if err = publicationRepository.Delete(publicationId); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func FindPublicationByUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
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
	publications, err := publicationRepository.FindPublicationByUser(userId)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, publications)

}

func LikePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationsId, err := strconv.ParseUint(params["publicationId"], 10, 64)
	if err != nil {
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
	if err = publicationRepository.LikePublication(publicationsId); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)

}
