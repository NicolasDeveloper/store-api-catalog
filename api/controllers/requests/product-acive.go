package requests

// ActiveProductRequest request type
// swagger:model
type ActiveProductRequest struct {
	// identification to active product
	// required: true
	ID string `json:"id"`
}
