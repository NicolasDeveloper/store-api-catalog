package requests

// ListCategorytRequest list products
// swagger:model
type ListCategorytRequest struct {
	PageSize         int `json:"page_size"`
	CurrentPageIndex int `json:"current_page_index"`
}
