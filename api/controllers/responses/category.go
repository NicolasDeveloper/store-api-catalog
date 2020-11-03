package responses

// CategoryResponse category view model
type CategoryResponse struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	ParentCategoryID string `json:"parent_category_id"`
}

// ResponseData response payload
//swagger:response categoryResponse
type swaggerCategoryResponse struct {
	//in: body
	Body struct {
		Success bool             `json:"success"`
		Data    CategoryResponse `json:"data"`
	}
}
