package requests

// RemoveSkuRequest update product variation
// swagger:model
type RemoveSkuRequest struct {
	// sku reference
	SkuID string `json:"sku_id"`
	// product reference
	ProductID string `json:"product_id"`
}
