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

// CreateSubscription is caled to create an subscription
func (server *Server) CreateSubscription(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	subscription := model.Subscription{}
	err = json.Unmarshal(body, &subscription)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = subscription.Validate("update")
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = subscription.Save(server.DB)

	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, subscription.ID))
	response.JSON(w, http.StatusCreated, subscription)
}

// GetSubscriptions retrieves all subscriptions
func (server *Server) GetSubscriptions(w http.ResponseWriter, r *http.Request) {
	var err error
	subscription := model.Subscription{}
	count, err := subscription.Count(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	data, err := subscription.FindAll(server.DB)
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

// GetSubscription loads an subscription by given ID
func (server *Server) GetSubscription(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := uuid.FromString(vars["id"])
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	subscription := model.Subscription{}
	err = subscription.FindByID(server.DB, uid)
	if err != nil {
		response.ERROR(w, http.StatusNotFound, err)
		return
	}
	response.JSON(w, http.StatusOK, subscription)
}

// UpdateSubscription updates existing subscription
func (server *Server) UpdateSubscription(w http.ResponseWriter, r *http.Request) {
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

	subscription := model.Subscription{}
	err = json.Unmarshal(body, &subscription)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = subscription.Validate("update")
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	subscription.ID = uid

	err = subscription.Update(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, subscription)
}

// DeleteSubscription deletes an subscription
func (server *Server) DeleteSubscription(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subscription := model.Subscription{}

	uid, err := uuid.FromString(vars["id"])
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = subscription.FindByID(server.DB, uid)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	err = subscription.Delete(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%s", uid))
	response.JSON(w, http.StatusNoContent, "")
}
