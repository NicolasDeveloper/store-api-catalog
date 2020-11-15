package responses

// ListPaginatorResponse list paginator response
type ListPaginatorResponse struct {
	NextPageSize int           `json:"next_page_size"`
	PrevPageSize int           `json:"prev_page_size"`
	CurrentPage  int           `json:"current_page"`
	PrevPage     int           `json:"prev_page"`
	Items        []interface{} `json:"items"`
}

// ResponseData response payload
//swagger:response listPaginatorResponse
type swaggerListPaginatorResponse struct {
	//in: body
	Body struct {
		Success bool                  `json:"success"`
		Data    ListPaginatorResponse `json:"data"`
	}
}
