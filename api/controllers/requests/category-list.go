package requests

// ListCategorytRequest list products
// swagger:model
type ListCategorytRequest struct {
	Size int `json:"size"`
	Page int `json:"page"`
}
