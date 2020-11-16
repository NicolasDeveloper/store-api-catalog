package responses

// ListPaginatorResponse list paginator response
type ListPaginatorResponse struct {
	CurrentPage  int `json:"current_page"`
	NextPage     int `json:"next_page"`
	PrevPage     int `json:"prev_page"`
	TotalPages   int `json:"total_pages"`
	NextPageSize int `json:"next_page_size"`
	PrevPageSize int `json:"prev_page_size"`
}
