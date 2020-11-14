package requests

// InactiveProductRequest request type
// swagger:model
type InactiveProductRequest struct {
	// product identification
	// required: true
	ID string `json:"id"`
}
