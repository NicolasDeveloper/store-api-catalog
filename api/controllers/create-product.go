package controllers

import (
	"net/http"

	"github.com/NicolasDeveloper/store-catalog-api/domain"
	"github.com/golobby/container"
)

//CreateProduct create product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	request := domain.CreateProductRequest{}
	GetContent(&request, r)

	product, err := domain.NewProduct(
		request.Name,
		request.Description,
	)

	var repository domain.ProductRepository
	container.Make(&repository)

	err = repository.Save(product)

	if err != nil {
		HandleError(err, w)
		return
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
