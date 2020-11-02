package controllers

import (
	"net/http"

	"github.com/NicolasDeveloper/store-catalog-api/domain"
	"github.com/golobby/container"
)

//CreateCategory create product
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	request := domain.CreateCategorytRequest{}

	GetContent(&request, r)

	category, err := domain.NewCategory(request.Name, request.Path)
	category.Link(request.ParentCategoryID)

	var repository domain.CategoryRepository
	container.Make(&repository)

	err = repository.Save(category)

	if err != nil {
		HandleError(err, w)
		return
	}

	resp := ResponseData{
		Success: true,
		Data: domain.CategoryResponse{
			ID:               category.ID,
			Name:             category.Name,
			ParentCategoryID: category.ParentCategoryID,
		},
	}

	SendJSON(w, resp, http.StatusCreated)
}
