package requests

// UpdateSkuRequest update product variation
// swagger:model
type UpdateSkuRequest struct {
	// sku reference
	SkuID string `json:"sku_id"`
	CreateSkuRequest
}
