package controller

import (
	"ais/config"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

const (
	internalServerErrorMessage = "Internal server error. We are sorry for inconvenience!"
	badFieldError              = "field name cannot be empty"
)

func respondWithError(ctx *gin.Context, code int, message string) error {
	return respondWithJSON(ctx, code, ErrorMessage{Message: message})
}

func respondWithJSON(ctx *gin.Context, code int, payload interface{}) error {
	response, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	w := ctx.Writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(response)
	if err != nil {
		return err
	}
	return nil
}

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(config.Salt)))
}
