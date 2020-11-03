package responses

//ResponseData object response
type ResponseData struct {
	Success bool `json:"success"`
	// Can be a ProductResponse or a CategoryResponse
	Data interface{} `json:"data"`
}

// ResponseData response payload
// swagger:response defaultResponse
type swaggResponseDataResp struct {
	// in:body
	Body ResponseData
}
