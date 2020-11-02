package controllers

import (
	"net/http"

	"github.com/NicolasDeveloper/store-catalog-api/domain"
	"github.com/golobby/container"
)

//UpdateCategory create product
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	request := domain.UpdateCategoryRequest{}

	GetContent(&request, r)

	var repository domain.CategoryRepository
	container.Make(&repository)

	category, err := repository.GetByID(request.ID)

	if err != nil {
		HandleError(err, w)
		return
	}

	category.SetName(request.Name)
	category.SetPath(request.Path)
	category.Link(request.ParentCategoryID)

	err = repository.Update(category)

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
