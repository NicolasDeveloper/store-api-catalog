package requests

// CreateCategorytRequest for creation of product
// swagger:model
type CreateCategorytRequest struct {
	// this field will appear in the catalog section title
	// required: true
	Name string `json:"name"`
	// this field will be used to find a specific category in URI
	// required: true
	// unique: true
	Path string `json:"path"`
	// this field will be used to link categories tree
	ParentCategoryID string `json:"parent_category_id"`
}
