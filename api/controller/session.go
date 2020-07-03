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

// CreateSession is caled to create an session
func (server *Server) CreateSession(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	session := model.Session{}
	err = json.Unmarshal(body, &session)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = session.Validate("update")
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = session.Save(server.DB)

	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, session.ID))
	response.JSON(w, http.StatusCreated, session)
}

// GetSessions retrieves all sessions
func (server *Server) GetSessions(w http.ResponseWriter, r *http.Request) {
	var err error
	session := model.Session{}
	count, err := session.Count(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	data, err := session.FindAll(server.DB)
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

// GetSession loads an session by given ID
func (server *Server) GetSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := uuid.FromString(vars["id"])
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	session := model.Session{}
	err = session.FindByID(server.DB, uid)
	if err != nil {
		response.ERROR(w, http.StatusNotFound, err)
		return
	}
	response.JSON(w, http.StatusOK, session)
}

// UpdateSession updates existing session
func (server *Server) UpdateSession(w http.ResponseWriter, r *http.Request) {
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

	session := model.Session{}
	err = json.Unmarshal(body, &session)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = session.Validate("update")
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	session.ID = uid

	err = session.Update(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, session)
}

// DeleteSession deletes an session
func (server *Server) DeleteSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	session := model.Session{}

	uid, err := uuid.FromString(vars["id"])
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = session.FindByID(server.DB, uid)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	err = session.Delete(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%s", uid))
	response.JSON(w, http.StatusNoContent, "")
}
