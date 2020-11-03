package requests

// CreateProductRequest request model
// swagger:model
type CreateProductRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Categories  []string `json:"categories"`
}
