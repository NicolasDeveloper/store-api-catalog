package requests

// UpdateProductRequest Update product
// swagger:model
type UpdateProductRequest struct {
	// category identification
	// required: true
	ID string `json:"id"`
	CreateProductRequest
}
