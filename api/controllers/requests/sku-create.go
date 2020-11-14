package requests

// CreateSkuRequest create product variation
// swagger:model
type CreateSkuRequest struct {
	// variation name
	Name string `json:"name"`
	// price of product
	Price float64 `json:"price"`
	// quantity of product
	Quantity int `json:"quantity"`
	// product reference
	ProductID string `json:"product_id"`
}
