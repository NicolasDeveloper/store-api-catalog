package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/NicolasDeveloper/store-catalog-api/api/controllers/requests"
	"github.com/NicolasDeveloper/store-catalog-api/domain"
	"github.com/golobby/container"

	"github.com/NicolasDeveloper/store-catalog-api/api/controllers/responses"
)

// LisCategories list categorires
// swagger:operation GET /categories/ categories listCategories
// ---
// summary: List category.
// description: list categorires
// parameters:
// - name: page_size
//   description: batch items size
//   required: true
//   in: query
// - name: current_page_index
//   description: current item index
//   required: true
//   in: query
// responses:
//   "200":
//     "$ref": "#/responses/categoryPaginatorResponse"
func LisCategories(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	currentPageIndex, err := strconv.ParseInt(query.Get("current_page_index"), 10, 64)

	if err != nil {
		HandleError(errors.New("Missing page index start"), w)
		return
	}

	pageSize, err := strconv.ParseInt(query.Get("page_size"), 10, 64)

	if err != nil {
		HandleError(errors.New("Missing page page size"), w)
		return
	}

	request := requests.ListCategorytRequest{
		CurrentPageIndex: int(currentPageIndex),
		PageSize:         int(pageSize),
	}

	var repository domain.CategoryRepository
	container.Make(&repository)

	total, err := repository.Total()

	if err != nil {
		HandleError(errors.New("Total not found"), w)
		return
	}

	paginator, err := domain.NewPaginator(request.CurrentPageIndex, request.PageSize, total)

	if err != nil {
		HandleError(err, w)
		return
	}

	categories, err := repository.ListCategories(paginator.GetSkip(), paginator.PageSize)

	if err != nil {
		HandleError(errors.New("Categories not found"), w)
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
				NextPage:     paginator.GetNextPage(),
				PrevPage:     paginator.GetPrevPage(),
				NextPageSize: paginator.GetNextPageIndex(),
				PrevPageSize: paginator.GetPrevPageIndex(),
				TotalPages:   paginator.GetTotalPageQuantity(),
			},
		},
	}

	SendJSON(w, resp, http.StatusCreated)
}
