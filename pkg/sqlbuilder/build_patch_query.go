package sqlbuilder

import (
	"errors"
	"fmt"

	"main.go/interfaces/http/dto"
)

func BuildPatchQuery(id int, product *dto.ProductPatchDTO) ([]string, []interface{}, int, error) {
	var (
		set  = []string{}
		args = []interface{}{}
		num  = 1
	)
	if product.Name != nil {
		set = append(set, fmt.Sprintf("name = $%d", num))
		args = append(args, product.Name)
		num++
	}
	if product.Category != nil {
		set = append(set, fmt.Sprintf("category = $%d", num))
		args = append(args, product.Category)
		num++
	}
	if product.Price != nil {
		set = append(set, fmt.Sprintf("price = $%d", num))
		args = append(args, product.Price)
		num++
	}
	if product.InStock != nil {
		set = append(set, fmt.Sprintf("in_stock = $%d", num))
		args = append(args, product.InStock)
		num++
	}
	if len(set) == 0 {
		return nil, nil, 0, errors.New("empty set")
	}
	return set, args, num, nil
}
