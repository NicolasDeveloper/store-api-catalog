package controllers

import (
	"net/http"

	"github.com/NicolasDeveloper/store-catalog-api/api/controllers/requests"
	"github.com/NicolasDeveloper/store-catalog-api/domain"
	"github.com/golobby/container"

	"github.com/NicolasDeveloper/store-catalog-api/api/controllers/responses"
)

// LisCategories list categorires
// swagger:operation GET /categories/ categories  LisCategories
// ---
// summary: List category.
// description: list categorires
// parameters:
// - name:  ListRequesCategoriest
//   description: request model to update category
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/ ListRequesCategoriest"
// responses:
//   "200":
//     "$ref": "#/responses/categoryResponse"
func LisCategories(w http.ResponseWriter, r *http.Request) {
	request := requests.ListCategorytRequest{}

	GetContent(&request, r)

	var repository domain.CategoryRepository
	container.Make(&repository)

	total, err := repository.Total()

	if err != nil {
		HandleError(err, w)
		return
	}

	paginator, err := domain.NewPaginator(request.CurrentPageIndex, request.PageSize, total)

	if err != nil {
		HandleError(err, w)
		return
	}

	categories, err := repository.ListCategories(paginator.GetSkip(), paginator.PageSize)

	if err != nil {
		HandleError(err, w)
		return
	}

	categoryResponse := make([]responses.CategoryResponse, 0)

	for _, category := range categories {
		categoryResponse = append(categoryResponse, responses.CategoryResponse{
			ID:               category.ID,
			Name:             category.Name,
			ParentCategoryID: category.ParentCategoryID,
		})
	}

	resp := responses.ResponseData{
		Success: true,
		Data: responses.CategoryPaginatorResponse{
			Items: categoryResponse,
			ListPaginatorResponse: responses.ListPaginatorResponse{
				CurrentPage:  paginator.GetCurrentPage(),
				PrevPage:     paginator.GetPrevPage(),
				NextPageSize: paginator.GetNextPageIndex(),
				PrevPageSize: paginator.GetPrevPageIndex(),
			},
		},
	}

	SendJSON(w, resp, http.StatusCreated)
}
