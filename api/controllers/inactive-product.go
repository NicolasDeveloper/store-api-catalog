package controllers

import (
	"net/http"

	"github.com/NicolasDeveloper/store-catalog-api/domain"
	"github.com/golobby/container"
)

//InactiveProduct Inactive product
func InactiveProduct(w http.ResponseWriter, r *http.Request) {
	request := domain.UpdateProductRequest{}
	GetContent(&request, r)

	var repository domain.ProductRepository
	container.Make(&repository)

	product, err := repository.GetByID(request.ID)

	product.Disable()

	err = repository.Update(product)

	if err != nil {
		HandleError(err, w)
	}

	resp := ResponseData{
		Success: true,
		Data: domain.ProductResponse{
			ID:          product.ID,
			Description: product.Description,
			Name:        product.Name,
		},
	}

	SendJSON(w, resp, http.StatusCreated)
}
