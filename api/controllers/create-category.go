package controllers

import (
	"net/http"

	"github.com/NicolasDeveloper/store-catalog-api/api/controllers/requests"
	"github.com/NicolasDeveloper/store-catalog-api/api/controllers/responses"
	"github.com/NicolasDeveloper/store-catalog-api/services"

	"github.com/NicolasDeveloper/store-catalog-api/domain"
	"github.com/golobby/container"
)

// CreateCategory create category
// swagger:operation POST /categories/ categories createCategory
// ---
// summary: Create category.
// description: Create category to link with a product
// parameters:
// - name: CreateCategorytRequest
//   description: request model to create category
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/CreateCategorytRequest"
// responses:
//   "200":
//     "$ref": "#/responses/categoryResponse"
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	request := requests.CreateCategorytRequest{}

	GetContent(&request, r)

	category, err := domain.NewCategory(request.Name, request.Path)
	category.Link(request.ParentCategoryID)

	var repository domain.CategoryRepository
	container.Make(&repository)

	category, err = services.LinkCategory(request.ParentCategoryID, category, repository)

	if err != nil {
		HandleError(err, w)
		return
	}

	err = repository.Save(category)

	if err != nil {
		HandleError(err, w)
		return
	}

	resp := responses.ResponseData{
		Success: true,
		Data: responses.CategoryResponse{
			ID:               category.ID,
			Name:             category.Name,
			ParentCategoryID: category.ParentCategoryID,
		},
	}

	SendJSON(w, resp, http.StatusCreated)
}
