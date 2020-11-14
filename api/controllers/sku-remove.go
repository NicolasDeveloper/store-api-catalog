package controllers

import (
	"net/http"

	"github.com/NicolasDeveloper/store-catalog-api/api/controllers/requests"
	"github.com/NicolasDeveloper/store-catalog-api/api/controllers/responses"
	"github.com/NicolasDeveloper/store-catalog-api/domain"
	"github.com/golobby/container"
)

// RemoveSku create sku
// swagger:operation PUT /products/skus/remove/ skus RemoveSku
// ---
// summary: Remove sku.
// description: Remove a product variation
// parameters:
// - name: RemoveSkuRequest
//   description: request model to create sku
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/RemoveSkuRequest"
// responses:
//   "200":
//     "$ref": "#/responses/productResponse"
func RemoveSku(w http.ResponseWriter, r *http.Request) {
	request := requests.UpdateSkuRequest{}
	GetContent(&request, r)

	var productRepository domain.ProductRepository
	container.Make(&productRepository)

	product, err := productRepository.GetByID(request.ProductID)

	if err != nil {
		HandleError(err, w)
		return
	}

	err = product.RemoveSku(request.SkuID)

	if err != nil {
		HandleError(err, w)
		return
	}

	err = productRepository.Update(product)

	if err != nil {
		HandleError(err, w)
		return
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
