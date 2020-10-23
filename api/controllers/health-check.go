package controllers

import (
	"net/http"
)

//HealthCheck create product
func HealthCheck(w http.ResponseWriter, r *http.Request) {

	resp := ResponseData{
		Success: true,
		Data: HelthCheckResponse{
			Name:    "Catalog API",
			Version: "1.0",
		},
	}

	SendJSON(w, resp, http.StatusNoContent)
}
