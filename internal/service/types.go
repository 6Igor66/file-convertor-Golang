package service

type PostResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   struct {
		ID      string `json:"id"`
		Minutes int    `json:"minutes"`
	} `json:"data"`
}

type GetFileResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   struct {
		ID      string `json:"id"`
		Encode  string `json:"encode"`
		Content string `json:"content"`
	} `json:"data"`
}
