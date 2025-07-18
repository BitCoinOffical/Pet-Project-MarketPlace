package dto

type ProductResponse struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	InStock  bool    `json:"in_stock"`
}
type ProductList struct {
	Items      []ProductResponse `json:"items"`
	TotalCount int               `json:"total_count"`
}
