package controllers

import (
	"net/http"

	"github.com/NicolasDeveloper/store-catalog-api/api/controllers/requests"
	"github.com/NicolasDeveloper/store-catalog-api/api/controllers/responses"
	"github.com/NicolasDeveloper/store-catalog-api/domain"
	"github.com/golobby/container"
)

//CreateProduct create product
// swagger:operation POST /products/ products createProduct
// ---
// summary: Create a product.
// description: Create a product in catalog
// parameters:
// - name: CreateProductRequest
//   description: request model to create product
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/CreateProductRequest"
// responses:
//   "200":
//     "$ref": "#/responses/productRes"
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	request := requests.CreateProductRequest{}

	GetContent(&request, r)

	var categoryRepository domain.CategoryRepository
	container.Make(&categoryRepository)

	var productRepository domain.ProductRepository
	container.Make(&productRepository)

	categories, err := categoryRepository.GetCategories(request.Categories)

	if err != nil {
		HandleError(err, w)
		return
	}

	product, err := domain.NewProduct(
		request.Name,
		request.Description,
		categories,
	)

	err = productRepository.Save(product)

	if err != nil {
		HandleError(err, w)
		return
	}

	resp := ResponseData{
		Success: true,
		Data: responses.ProductResponse{
			ID:          product.ID,
			Description: product.Description,
			Name:        product.Name,
		},
	}

	SendJSON(w, resp, http.StatusCreated)
}
