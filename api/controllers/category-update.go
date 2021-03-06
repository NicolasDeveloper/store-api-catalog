package controllers

import (
	"net/http"

	"github.com/NicolasDeveloper/store-catalog-api/api/controllers/requests"
	"github.com/NicolasDeveloper/store-catalog-api/api/controllers/responses"
	"github.com/NicolasDeveloper/store-catalog-api/domain"
	"github.com/NicolasDeveloper/store-catalog-api/services"
	"github.com/golobby/container"
)

//UpdateCategory update category
// swagger:operation PUT /categories/ categories updateCategory
// ---
// summary: Update category.
// description: Update category to link with a product
// parameters:
// - name: UpdateCategoryRequest
//   description: request model to update category
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/UpdateCategoryRequest"
// responses:
//   "200":
//     "$ref": "#/responses/categoryResponse"
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	request := requests.UpdateCategoryRequest{}

	GetContent(&request, r)

	var repository domain.CategoryRepository
	container.Make(&repository)

	category, err := repository.GetByID(request.ID)

	if err != nil {
		HandleError(err, w)
		return
	}

	category.SetName(request.Name)
	category.SetPath(request.Path)

	category, err = services.LinkCategory(request.ParentCategoryID, category, repository)

	if err != nil {
		HandleError(err, w)
		return
	}

	err = repository.Update(category)

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
