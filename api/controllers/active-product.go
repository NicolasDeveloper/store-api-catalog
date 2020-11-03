package controllers

import (
	"net/http"

	"github.com/NicolasDeveloper/store-catalog-api/api/controllers/requests"
	"github.com/NicolasDeveloper/store-catalog-api/api/controllers/responses"
	"github.com/NicolasDeveloper/store-catalog-api/domain"
	"github.com/golobby/container"
)

// ActiveProduct create product
// swagger:operation POST /products/active/ products activeProduct
// ---
// summary: Active a product.
// description: If a product its disabled this enpoint can active again
// parameters:
// - name: ActiveProductRequest
//   description: request model to create category
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/ActiveProductRequest"
// responses:
//   "200":
//     "$ref": "#/responses/productResponse"
func ActiveProduct(w http.ResponseWriter, r *http.Request) {
	request := requests.UpdateProductRequest{}
	GetContent(&request, r)

	var repository domain.ProductRepository
	container.Make(&repository)

	product, err := repository.GetByID(request.ID)

	product.Enable()

	err = repository.Update(product)

	if err != nil {
		HandleError(err, w)
	}

	resp := responses.ResponseData{
		Success: true,
		Data: responses.ProductResponse{
			ID:          product.ID,
			Description: product.Description,
			Name:        product.Name,
		},
	}

	SendJSON(w, resp, http.StatusCreated)
}
