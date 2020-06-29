package controller

import (
	"net/http"

	"github.com/dzahariev/e2e-rest/api/response"
)

// Home is an API root route controller
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, "Welcome!")
}
