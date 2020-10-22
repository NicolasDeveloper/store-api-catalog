package domain

// UpdateProductRequest Update product
type UpdateProductRequest struct {
	ID string `json:"id"`
	CreateProductRequest
}
