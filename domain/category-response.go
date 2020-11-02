package domain

//CategoryResponse category view model
type CategoryResponse struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	ParentCategoryID string `json:"parent_category_id"`
}
