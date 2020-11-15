package web

type ResponseData struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
