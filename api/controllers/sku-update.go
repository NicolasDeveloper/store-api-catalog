package controllers

import (
	"net/http"

	"github.com/NicolasDeveloper/store-catalog-api/api/controllers/requests"
	"github.com/NicolasDeveloper/store-catalog-api/api/controllers/responses"
	"github.com/NicolasDeveloper/store-catalog-api/domain"
	"github.com/golobby/container"
)

//UpdateSku update sku
// swagger:operation PUT /products/skus/ skus updateSku
// ---
// summary: Update sku
// description: Update a product variation
// parameters:
// - name: UpdateSkuRequest
//   description: request model to update sku
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/UpdateSkuRequest"
// responses:
//   "200":
//     "$ref": "#/responses/productResponse"
func UpdateSku(w http.ResponseWriter, r *http.Request) {
	request := requests.UpdateSkuRequest{}
	GetContent(&request, r)

	var productRepository domain.ProductRepository
	container.Make(&productRepository)

	product, err := productRepository.GetByID(request.ProductID)

	if err != nil {
		HandleError(err, w)
		return
	}

	sku, err := product.FindSku(request.SkuID)

	if err != nil {
		HandleError(err, w)
		return
	}

	product.UpdateSku(sku)

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
