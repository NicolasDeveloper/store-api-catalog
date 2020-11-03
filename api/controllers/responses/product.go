package responses

// ProductResponse product view model
type ProductResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ResponseData response payload
//swagger:response productResponse
type swaggerproductResponse struct {
	//in: body
	Body struct {
		Success bool            `json:"success"`
		Data    ProductResponse `json:"data"`
	}
}
