package responses

// ProductResponse product view model
type ProductResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ProductResponse response payload
// swagger:response productRes
type swaggProductRes struct {
	// in:body
	Body ProductResponse
}
