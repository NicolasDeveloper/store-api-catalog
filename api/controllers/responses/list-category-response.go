package responses

// ListPaginatorResponse list paginator response
type ListPaginatorResponse struct {
	NextPageSize int `json:"next_page_size"`
	PrevPageSize int `json:"prev_page_size"`
	CurrentPage  int `json:"current_page"`
	PrevPage     int `json:"prev_page"`
}
