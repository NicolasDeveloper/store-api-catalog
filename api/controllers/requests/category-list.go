package requests

// ListCategorytRequest list products
// swagger:model listCategories
type ListCategorytRequest struct {
	// PageSize page size
	PageSize int `json:"page_size"`
	// CurrentPageIndex current index size
	CurrentPageIndex int `json:"current_page_index"`
}
