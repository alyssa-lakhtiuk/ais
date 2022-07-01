package controller

import (
	"ais/config"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	internalServerErrorMessage = "Internal server error. We are sorry for inconvenience!"
	badFieldError              = "field name cannot be empty"
)

func respondWithError(w http.ResponseWriter, code int, message string) error {
	return respondWithJSON(w, code, ErrorMessage{Message: message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) error {
	response, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
	return nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(config.Salt)))
}
