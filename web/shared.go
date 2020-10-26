package web

type ResponseData struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

type Metadata struct {
	Created string `json:"created"`
	Edited  string `json:"edited"`
	Url     string `json:"url"`
}
