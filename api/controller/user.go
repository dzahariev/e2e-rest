package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dzahariev/e2e-rest/api/middleware"
	"github.com/dzahariev/e2e-rest/api/model"
	"github.com/dzahariev/e2e-rest/api/response"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
)

// CreateUser is caled to create an user
func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := model.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = user.Validate("update")
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = user.Save(server.DB)

	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, user.ID))
	response.JSON(w, http.StatusCreated, user)
}

// GetUsers retrieves all users
func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	var err error
	user := model.User{}
	count, err := user.Count(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	data, err := user.FindAll(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	list := model.List{
		Count: count,
		Data:  *data,
	}

	response.JSON(w, http.StatusOK, list)
}

// GetUser loads an user by given ID
func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := uuid.FromString(vars["id"])
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	user := model.User{}
	err = user.FindByID(server.DB, uid)
	if err != nil {
		response.ERROR(w, http.StatusNotFound, err)
		return
	}
	response.JSON(w, http.StatusOK, user)
}

// UpdateUser updates existing user
func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := uuid.FromString(vars["id"])
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := model.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	userIDFromContext := r.Context().Value(middleware.KeyUserID)
	if userIDFromContext != uid {
		response.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	err = user.Validate("update")
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.ID = uid

	err = user.Update(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, user)
}

// DeleteUser deletes an user
func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := model.User{}

	uid, err := uuid.FromString(vars["id"])
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	userIDFromContext := r.Context().Value(middleware.KeyUserID)
	if userIDFromContext != uid {
		response.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	err = user.FindByID(server.DB, uid)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	err = user.Delete(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%s", uid))
	response.JSON(w, http.StatusNoContent, "")
}
