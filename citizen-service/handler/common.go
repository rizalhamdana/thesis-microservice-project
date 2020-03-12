package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/rizalhamdana/citizen-service/model"
)

//TokenClaims is a struct for token payload
type TokenClaims struct {
	Nik  string `json:"nik"`
	Name string `json:"name"`
	jwt.StandardClaims
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}

func createToken(citizen *model.Citizen) (string, error) {
	var jwtKey = []byte("citizen_service_secret_key")
	expirationTime := time.Now().Add(1440 * time.Minute)

	claims := &TokenClaims{
		Nik:  citizen.NIK,
		Name: citizen.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil

}
