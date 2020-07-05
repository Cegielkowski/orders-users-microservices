package controllers

import (
	"encoding/json"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/user-api/pkg/errors"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/user-api/pkg/responses"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/user-api/pkg/services/users"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/user-api/pkg/structs"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// GetUserController Get the user information of the requested id.
func GetUserController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	userID, err := strconv.ParseUint(vars["userID"], 10, 32)
	if err != nil {
		errors.Log(err)
		responses.Finish(w, &responses.HttpBadRequest{})
		return
	}

	ret, err := users.Get(uint32(userID))
	if err != nil {
		responses.Finish(w, errors.GetHttpResponse(err))
		return
	}

	err = json.NewEncoder(w).Encode(ret)
	if err != nil {
		responses.Finish(w, errors.GetHttpResponse(err))
	}
}

// ListUserController List the users.
func ListUserController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ret, err := users.List()
	if err != nil {
		responses.Finish(w, errors.GetHttpResponse(err))
		return
	}

	err = json.NewEncoder(w).Encode(ret)
	if err != nil {
		responses.Finish(w, errors.GetHttpResponse(err))
	}
}

// PostUserController Creates a new user.
func PostUserController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user structs.PostUserRequest
	err := json.NewDecoder(r.Body).Decode(&user)

	defer func() { r.Body.Close() }()

	if err != nil {
		errors.Log(err)
		responses.Finish(w, &responses.HttpBadRequest{})
		return
	}

	created, err := users.Insert(user.Email, user.Cpf, user.Phone, user.Name)
	if err != nil {
		errors.Log(err)
		responses.Finish(w, errors.GetHttpResponse(err))
		return
	}

	responses.Finish(w, &responses.HttpCreated{
		Content: created,
	})
}

// PutUserController Updates a user.
func PutUserController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	userID, err := strconv.ParseUint(vars["userID"], 10, 32)
	if err != nil {
		errors.Log(err)
		responses.Finish(w, &responses.HttpBadRequest{})
		return
	}

	var user structs.PostUserRequest
	err = json.NewDecoder(r.Body).Decode(&user)

	defer func() { r.Body.Close() }()

	if err != nil {
		errors.Log(err)
		responses.Finish(w, &responses.HttpBadRequest{})
		return
	}

	updated, err := users.Update(uint32(userID), user)
	if err != nil {
		errors.Log(err)
		responses.Finish(w, errors.GetHttpResponse(err))
		return
	}

	responses.Finish(w, &responses.HttpSuccessContent{
		Content: updated,
	})
}

// DeleteUserController Delete a user.
func DeleteUserController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	userID, err := strconv.ParseUint(vars["userID"], 10, 32)
	if err != nil {
		errors.Log(err)
		responses.Finish(w, &responses.HttpBadRequest{})
		return
	}

	err = users.Delete(uint32(userID))
	if err != nil {
		responses.Finish(w, errors.GetHttpResponse(err))
		return
	}

	responses.Finish(w, &responses.HttpSuccessContent{})
}