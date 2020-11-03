package requests

// UpdateProductRequest Update product
// swagger:model
type UpdateProductRequest struct {
	ID string `json:"id"`
	CreateProductRequest
}
