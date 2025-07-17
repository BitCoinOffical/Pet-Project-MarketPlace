package products

import "time"

type Product struct {
	ID        int
	Name      string
	Category  string
	Price     float64
	InStock   bool
	CreatedAt time.Time
}
