package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// SendJSON marshals v to a json struct and sends appropriate headers to w
func SendJSON(w http.ResponseWriter, v interface{}, code int) {
	w.Header().Add("Content-Type", "application/json")

	b, err := json.Marshal(v)

	if err != nil {
		log.Print(fmt.Sprintf("Error while encoding JSON: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error": "Internal server error"}`)
	} else {
		w.WriteHeader(code)
		io.WriteString(w, string(b))
	}
}

// GetContent of the request inside given struct
func GetContent(v interface{}, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(v)
	if err != nil {
		return err
	}

	return nil
}

// HandleError write error on response and return false if there is no error
func HandleError(err error, w http.ResponseWriter) bool {
	if err == nil {
		return false
	}

	msg := map[string]string{
		"message": "An error occured",
	}

	SendJSON(w, &msg, http.StatusInternalServerError)
	return true
}
