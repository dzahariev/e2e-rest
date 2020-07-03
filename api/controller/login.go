package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dzahariev/e2e-rest/api/auth"
	"github.com/dzahariev/e2e-rest/api/model"
	"github.com/dzahariev/e2e-rest/api/response"
	"golang.org/x/crypto/bcrypt"
)

// LogIn returns a token for user
func (server *Server) LogIn(w http.ResponseWriter, r *http.Request) {
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

	err = user.Validate("login")
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.GetTokenForUser(user.Email, user.Password)
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			response.ERROR(w, http.StatusUnauthorized, err)
		} else {
			response.ERROR(w, http.StatusUnprocessableEntity, err)
		}
		return
	}
	response.JSON(w, http.StatusOK, token)
}

// GetTokenForUser returns a token for the user
func (server *Server) GetTokenForUser(email, password string) (token string, err error) {
	user := model.User{}
	err = server.DB.Model(model.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = model.VerifyPassword(user.Password, password)
	if err != nil {
		return "", err
	}
	return auth.CreateToken(user.ID)
}
