package responses

// ProductResponse product view model
// swagger:response productRes
type ProductResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ProductResponse response payload
// swagger:response productRes
type swaggProductResp struct {
	// in:body
	Body ProductResponse
}
