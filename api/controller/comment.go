package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dzahariev/e2e-rest/api/model"
	"github.com/dzahariev/e2e-rest/api/response"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
)

// CreateComment is caled to create an comment
func (server *Server) CreateComment(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	comment := model.Comment{}
	err = json.Unmarshal(body, &comment)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = comment.Validate("update")
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = comment.Save(server.DB)

	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, comment.ID))
	response.JSON(w, http.StatusCreated, comment)
}

// GetComments retrieves all comments
func (server *Server) GetComments(w http.ResponseWriter, r *http.Request) {
	var err error
	comment := model.Comment{}
	count, err := comment.Count(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	data, err := comment.FindAll(server.DB)
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

// GetComment loads an comment by given ID
func (server *Server) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := uuid.FromString(vars["id"])
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	comment := model.Comment{}
	err = comment.FindByID(server.DB, uid)
	if err != nil {
		response.ERROR(w, http.StatusNotFound, err)
		return
	}
	response.JSON(w, http.StatusOK, comment)
}

// UpdateComment updates existing comment
func (server *Server) UpdateComment(w http.ResponseWriter, r *http.Request) {
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

	comment := model.Comment{}
	err = json.Unmarshal(body, &comment)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = comment.Validate("update")
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	comment.ID = uid

	err = comment.Update(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, comment)
}

// DeleteComment deletes an comment
func (server *Server) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	comment := model.Comment{}

	uid, err := uuid.FromString(vars["id"])
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = comment.FindByID(server.DB, uid)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	err = comment.Delete(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%s", uid))
	response.JSON(w, http.StatusNoContent, "")
}
