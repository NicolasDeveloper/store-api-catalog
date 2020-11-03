package controllers

import (
	"net/http"

	"github.com/NicolasDeveloper/store-catalog-api/api/controllers/responses"
)

// HealthCheck create product
// swagger:operation GET /health-check/ health-check healthCheck
// ---
// summary: Health Check.
// description: Consult if api its alive
// responses:
//   "200":
//     "$ref": "#/responses/helthCheckResponse"
func HealthCheck(w http.ResponseWriter, r *http.Request) {

	resp := responses.ResponseData{
		Success: true,
		Data: responses.HelthCheckResponse{
			Name:    "Catalog API",
			Version: "1.0",
		},
	}

	SendJSON(w, resp, http.StatusOK)
}
