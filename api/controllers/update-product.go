package controllers

import (
	"net/http"

	"github.com/NicolasDeveloper/store-catalog-api/api/controllers/requests"
	"github.com/NicolasDeveloper/store-catalog-api/api/controllers/responses"
	"github.com/NicolasDeveloper/store-catalog-api/domain"
	"github.com/golobby/container"
)

//UpdateProduct Update product
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	request := requests.UpdateProductRequest{}
	GetContent(&request, r)

	var repository domain.ProductRepository
	container.Make(&repository)

	product, err := repository.GetByID(request.ID)

	product.SetDescription(request.Description)
	product.SetName(request.Name)

	err = repository.Update(product)

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
