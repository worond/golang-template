package model

type Pagination struct {
	Sort   string `json:"sort"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Filter string `json:"filter"`
}
