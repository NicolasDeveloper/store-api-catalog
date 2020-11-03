package responses

//HelthCheckResponse response
type HelthCheckResponse struct {
	Version string `json:"version"`
	Name    string `json:"name"`
}

// HelthCheckResponse payload response
// swagger:response helthCheckResponse
type swaggerHelthCheckResponse struct {
	Body HelthCheckResponse
}
