package requests

// CreateCategorytRequest for creation of product
// swagger:model
type CreateCategorytRequest struct {
	Name             string `json:"name"`
	Path             string `json:"path"`
	ParentCategoryID string `json:"parent_category_id"`
}
