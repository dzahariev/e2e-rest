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

// CreateEvent is caled to create an event
func (server *Server) CreateEvent(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	event := model.Event{}
	err = json.Unmarshal(body, &event)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = event.Validate("update")
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = event.Save(server.DB)

	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, event.ID))
	response.JSON(w, http.StatusCreated, event)
}

// GetEvents retrieves all events
func (server *Server) GetEvents(w http.ResponseWriter, r *http.Request) {
	var err error
	event := model.Event{}
	count, err := event.Count(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	data, err := event.FindAll(server.DB)
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

// GetEvent loads an event by given ID
func (server *Server) GetEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := uuid.FromString(vars["id"])
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	event := model.Event{}
	err = event.FindByID(server.DB, uid)
	if err != nil {
		response.ERROR(w, http.StatusNotFound, err)
		return
	}
	response.JSON(w, http.StatusOK, event)
}

// UpdateEvent updates existing event
func (server *Server) UpdateEvent(w http.ResponseWriter, r *http.Request) {
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

	event := model.Event{}
	err = json.Unmarshal(body, &event)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = event.Validate("update")
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	event.ID = uid

	err = event.Update(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, event)
}

// DeleteEvent deletes an event
func (server *Server) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	event := model.Event{}

	uid, err := uuid.FromString(vars["id"])
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = event.FindByID(server.DB, uid)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	err = event.Delete(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%s", uid))
	response.JSON(w, http.StatusNoContent, "")
}
