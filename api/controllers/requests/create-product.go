package requests

// CreateProductRequest request model
type CreateProductRequest struct {
	// the name for this product
	// required: true
	Name string `json:"name"`
	// the description for this product
	// required: true
	Description string `json:"description"`
	// the categories for this product
	// required: true
	Categories []string `json:"categories"`
}

// CreateProductRequest response payload
// swagger:response productReq
type swaggProductReq struct {
	// in:body
	Body CreateProductRequest
}
