package controllers

import (
	"net/http"

	"github.com/NicolasDeveloper/store-catalog-api/api/controllers/requests"
	"github.com/NicolasDeveloper/store-catalog-api/api/controllers/responses"
	"github.com/NicolasDeveloper/store-catalog-api/domain"
	"github.com/golobby/container"
)

//InactiveProduct Inactive product
// swagger:operation PUT /products/active/ products activeProduct
// ---
// summary: Active product.
// description: If product its active this enpoint can inactive
// parameters:
// - name: InactiveProductRequest
//   description: request model to inactive product
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/InactiveProductRequest"
// responses:
//   "200":
//     "$ref": "#/responses/productResponse"
func InactiveProduct(w http.ResponseWriter, r *http.Request) {
	request := requests.UpdateProductRequest{}
	GetContent(&request, r)

	var repository domain.ProductRepository
	container.Make(&repository)

	product, err := repository.GetByID(request.ID)

	product.Disable()

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
