package requests

// UpdateCategoryRequest Update category
// swagger:model
type UpdateCategoryRequest struct {
	// category identification
	// required: true
	ID string `json:"id"`
	CreateCategorytRequest
}
