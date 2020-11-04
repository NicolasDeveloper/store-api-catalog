package requests

// CreateProductRequest request model
// swagger:model
type CreateProductRequest struct {
	// product name
	// required: true
	Name string `json:"name"`
	// product description
	// required: true
	Description string `json:"description"`
	// this field will be used to link product with categories
	// required: true
	Categories []string `json:"categories"`
}
