package requests

//CreateCategorytRequest for creation of product
type CreateCategorytRequest struct {
	Name             string `json:"name"`
	Path             string `json:"path"`
	ParentCategoryID string `json:"parent_category_id"`
}

// CreateCategorytRequest response payload
// swagger:response categoryReq
type swaggCategoryReq struct {
	// in:body
	Body CreateCategorytRequest
}
