package requests

// CreateSkuRequest create product variation
// swagger:model request model to create a product variation
type CreateSkuRequest struct {
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
	ProductID string  `json:"product_id"`
}
