package domain

//CreateProductRequest for creation of product
type CreateProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
