package controllers

import (
	"net/http"

	"github.com/NicolasDeveloper/store-catalog-api/api/controllers/requests"
	"github.com/NicolasDeveloper/store-catalog-api/api/controllers/responses"
	"github.com/NicolasDeveloper/store-catalog-api/domain"
	"github.com/golobby/container"
)

//CreateSku create sku
// swagger:operation POST /products/skus/ skus createSku
// ---
// summary: Create sku.
// description: Create a product variation
// parameters:
// - name: CreateSkuRequest
//   description: request model to create sku
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/CreateSkuRequest"
// responses:
//   "200":
//     "$ref": "#/responses/productResponse"
func CreateSku(w http.ResponseWriter, r *http.Request) {
	request := requests.CreateSkuRequest{}
	GetContent(&request, r)

	var productRepository domain.ProductRepository
	container.Make(&productRepository)

	product, err := productRepository.GetByID(request.ProductID)

	if err != nil {
		HandleError(err, w)
	}

	product.AddSku(domain.NewSku(request.Name, request.Price, request.Quantity))

	err = productRepository.Update(product)

	if err != nil {
		HandleError(err, w)
	}

	productResponse := responses.ProductResponse{
		ID:          product.ID,
		Description: product.Description,
		Name:        product.Name,
	}

	productResponse.AddSkus(product.Skus)

	resp := responses.ResponseData{
		Success: true,
		Data:    productResponse,
	}

	SendJSON(w, resp, http.StatusCreated)
}
