package responses

// CategoryPaginatorResponse list paginator response
type CategoryPaginatorResponse struct {
	Items                 []CategoryResponse `json:"items"`
	ListPaginatorResponse `json:",inline"`
}

// ResponseData response payload
//swagger:response CategoryPaginatorResponse
type swaggerCategoryPaginatorResponse struct {
	//in: body
	Body struct {
		Success bool                      `json:"success"`
		Data    CategoryPaginatorResponse `json:"data"`
	}
}
