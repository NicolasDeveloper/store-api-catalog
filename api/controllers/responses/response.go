package responses

//ResponseData object response
type ResponseData struct {
	// this field indicates if everything has gone with success
	Success bool `json:"success"`
	// generic response data
	Data interface{} `json:"data"`
}

// ResponseData response payload
// swagger:response defaultResponse
type swaggResponseDataResp struct {
	// in:body
	Body ResponseData
}
