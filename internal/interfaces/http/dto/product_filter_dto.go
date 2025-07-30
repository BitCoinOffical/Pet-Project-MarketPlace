package dto

type ProductFilterDTO struct {
	Category *string
	MinPrice *float64
	MaxPrice *float64
	Search   *string
	Page     int
	Limit    int
}
