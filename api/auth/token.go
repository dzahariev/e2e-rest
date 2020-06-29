package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
)

// CreateToken creates a token
func CreateToken(userID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID.String()
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func extractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

// ExtractJWTToken extracts and returns a jwt token from request
func ExtractJWTToken(r *http.Request) (*jwt.Token, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil

}

// ValidateToken validates the token in request
func ValidateToken(token *jwt.Token) error {
	if !token.Valid {
		return fmt.Errorf("token is not valid")
	}

	return nil
}

// ExtractUserID extract the user from token in request
func ExtractUserID(token *jwt.Token) (uuid.UUID, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		uid, err := uuid.FromString(fmt.Sprintf("%s", claims["user_id"]))
		if err != nil {
			return uuid.Nil, fmt.Errorf("failed to parse UUID %v: %v", claims["user_id"], err)
		}

		return uid, nil
	}
	return uuid.Nil, fmt.Errorf("failed to extract UUID")
}
