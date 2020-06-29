package middleware

import (
	"errors"
	"log"
	"net/http"

	"github.com/dzahariev/e2e-rest/api/auth"
	"github.com/dzahariev/e2e-rest/api/response"
	"golang.org/x/net/context"
)

// Key is a named type for keys
type Key int

const (
	// KeyToken is used to store token in context
	KeyToken Key = iota

	// KeyUserID is used to store user ID in context
	KeyUserID Key = iota
)

// ContentTypeJSON set the content type to JSON
func ContentTypeJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

// CheckAuthentication check the auhtorisation
func CheckAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := auth.ExtractJWTToken(r)
		if err != nil {
			log.Println("error when extracting token:", err)
			response.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		err = auth.ValidateToken(token)
		if err != nil {
			log.Println("error when validating token:", err)
			response.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		// Adding Token to current request context
		newContext := context.WithValue(r.Context(), KeyToken, token)

		userID, err := auth.ExtractUserID(token)
		if err != nil {
			log.Println("error when extracting user from token:", err)
			response.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		// Adding UserID to current request context
		newContext = context.WithValue(newContext, KeyUserID, userID)
		r = r.WithContext(newContext)
		next(w, r)
	}
}
