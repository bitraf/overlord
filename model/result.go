package model

type PagedResult struct {
	Offset int         `json:"offset"`
	Limit  int         `json:"limit"`
	Count  int         `json:"count"`
	Total  int         `json:"total"`
	Result interface{} `json:"result"`
}
