package controllers

//ResponseData object response
type ResponseData struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}
