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
		Data: domain.ProductResponse{
			ID:          product.ID,
			Description: product.Description,
			Name:        product.Name,
		},
	}

	SendJSON(w, resp, http.StatusCreated)
}
