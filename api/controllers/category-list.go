package controllers

import (
	"net/http"

	"github.com/NicolasDeveloper/store-catalog-api/api/controllers/requests"
	"github.com/NicolasDeveloper/store-catalog-api/domain"
	"github.com/golobby/container"

	"github.com/NicolasDeveloper/store-catalog-api/api/controllers/responses"
)

// LisCategories list categorires
// swagger:operation GET /categories/ categories  LisCategories
// ---
// summary: List category.
// description: list categorires
// parameters:
// - name:  ListRequesCategoriest
//   description: request model to update category
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/ ListRequesCategoriest"
// responses:
//   "200":
//     "$ref": "#/responses/categoryResponse"
func LisCategories(w http.ResponseWriter, r *http.Request) {
	request := requests.ListCategorytRequest{}

	GetContent(&request, r)

	var repository domain.CategoryRepository
	container.Make(&repository)

	// filter, _ := services.CalculatePageSize(
	// 	request.Page,
	// 	request.Size,
	// 	repository,
	// )

	resp := responses.ResponseData{
		Success: true,
		Data:    nil,
	}

	SendJSON(w, resp, http.StatusCreated)
}
