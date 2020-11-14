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
// summary: Create product.
// description: Create product in catalog
// parameters:
// - name: CreateProductRequest
//   description: request model to create product
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/CreateProductRequest"
// responses:
//   "200":
//     "$ref": "#/responses/productResponse"
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

	categoryIDs := make([]string, 0)

	for _, category := range categories {
		categoryIDs = append(categoryIDs, category.ID)
	}

	product, err := domain.NewProduct(
		request.Name,
		request.Description,
		categoryIDs,
	)

	err = productRepository.Save(product)

	if err != nil {
		HandleError(err, w)
		return
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
